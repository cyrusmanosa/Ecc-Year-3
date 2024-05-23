package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type TempData struct {
	Temp      float32
	EntryTime time.Time
}

// TempDataの文字列化
func (td TempData) String() string {
	return fmt.Sprint("温度:", td.Temp, "℃ 取得時間:", td.EntryTime)
}

// 温度とその取得時間を持つスライス
var tempDataSlice = make([]TempData, 0)
var tempDataMutex sync.Mutex // グローバル変数へのロック制御行うためのmutex

var wg sync.WaitGroup // 温度センサの設定を待つために使用

var ledBlinkChan chan int = make(chan int) // LED点滅用チャネル

// 取得間隔の変数およびロック用Mutex
var intervalTime = 10
var intervalTimeMutex sync.Mutex

// 取得間隔の読み込み用関数
func getIntervalTime() int {
	// アクセスの際にmutexによりロックをする
	intervalTimeMutex.Lock()
	defer intervalTimeMutex.Unlock()
	return intervalTime
}

// 取得間隔の書き込み用関数
func setIntervalTime(interval int) {
	// 書き込みについても同様にロックをする
	intervalTimeMutex.Lock()
	defer intervalTimeMutex.Unlock()
	intervalTime = interval
}

// main関数では温度取得の開始およびWebサーバの起動を行う
func main() {
	wg.Add(1)
	go getTemp() // 温度データ取得のゴルーチンを開始

	wg.Add(1)
	go setLed() // LED点滅のゴルーチンを開始

	wg.Wait() // 温度センサおよびLEDの設定を待つ

	go tempAlert() // 気温アラート機能のゴルーチン開始を追加

	// Webサーバの設定と起動
	http.HandleFunc("/json", getJsonData)
	http.HandleFunc("/setting_interval", settingInterval)
	http.ListenAndServe(":4000", nil)
}

func getTemp() {
	// 温度センサADT7410を1 SPSモード、16ビット精度に設定を変更
	bus := embd.NewI2CBus(1)
	bus.WriteByteToReg(0x48, 0x03, 0xC0)

	getTempData := func() {
		ledBlinkChan <- 1 // LEDを点滅させるようsetLedに通知

		data16, _ := bus.ReadWordFromReg(0x48, 0x00) // 温度データを取得
		t := float32(data16) / 128.0                 // セ氏温度に変換

		// グローバル変数へのアクセス時はロックを取得 … ③
		tempDataMutex.Lock()         // 変数へアクセスする際にロックする
		defer tempDataMutex.Unlock() // 変数へのアクセスが終わったらロックを解除する

		tempData := TempData{t, time.Now()}

		// 下のコメントアウトを外せば取得のたびにコンソールで確認できる
		// fmt.Println(tempData)

		// スライスの最後に取得データを追加
		tempDataSlice = append(tempDataSlice, tempData)

		// 最新の10000件を残すように変更
		// dataLimit := 100
		dataLimit := 10000 // 1万件に変更
		if len(tempDataSlice) > dataLimit {
			tempDataSlice = tempDataSlice[len(tempDataSlice)-dataLimit:]
		}
	}

	tick := time.Tick(time.Second) // 1秒毎に受信するチャネルを生成
	<-tick                         // ここでは温度センサの取得完了を待つ … ②

	getTempData() // 温度データの取得

	wg.Done() // 温度センサ設定・取得の終了をメインゴルーチンに伝える … ①

	// 10秒ごとに温度データを取得し、取得時刻と温度を変数に格納
	// intervalTime := 10 //削除
	for {
		i := 0
		// for i < intervalTime{ // 削除
		for i < getIntervalTime() { // IntervalTimeに代えてgetIntervalTime()を利用する
			select {
			case <-tick:
				i++
			}
		}
		getTempData()
	}
}

// 取得した温度データをJSON形式に変換して出力
func getJsonData(w http.ResponseWriter, r *http.Request) {
	// グローバル変数へのアクセス時はロックを取得
	tempDataMutex.Lock()
	defer tempDataMutex.Unlock()

	jsonData, err := json.Marshal(tempDataSlice)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprint(w, string(jsonData))
}

func setLed() {
	// LEDをつなげたGPIOの設定
	embd.InitGPIO()
	pin, err := embd.NewDigitalPin(4)
	pin.SetDirection(embd.Out)

	if err != nil {
		fmt.Println(err)
	}

	// CTRL-Cでの終了時にGPIO設定をCloseする
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		pin.Close()
		os.Exit(0)
	}()

	wg.Done() // LEDの設定完了をメインゴルーチンに伝える

	for {
		<-ledBlinkChan
		if err := pin.Write(embd.High); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 200)
		if err := pin.Write(embd.Low); err != nil {
			panic(err)
		}
	}
}

// テンプレート用構造体
type ViewData struct {
	IntervalList     []string // 取得間隔設定のリスト
	SelectedInterval string   // 現在選択中の取得間隔
}

// テンプレートの設定
var tmpl *template.Template = template.Must(template.ParseFiles("setting.tmpl"))

func (v *ViewData) Execute(w http.ResponseWriter) {
	if err := tmpl.Execute(w, v); err != nil {
		panic(err)
	}
}

// 取得間隔の設定
func settingInterval(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // POSTの場合には待機時間を更新する
		formValue := r.FormValue("interval")
		intervalInt, err := strconv.Atoi(formValue)
		if err == nil {
			setIntervalTime(intervalInt) // 更新用関数を使用
		}
	}

	// 表示用データを設定して、クライアントに送信する
	v := ViewData{[]string{"1", "2", "10"}, strconv.Itoa(getIntervalTime())}
	v.Execute(w)
}

// 10分ごとにdiffTempAlertをコールする
func tempAlert() {
	diffTempAlert()
	tick := time.Tick(time.Second * 60) // 10分ごとに受信するチャネルを生成
	for {
		i := 0
		for i < getIntervalTime() {
			select {
			case <-tick:
				i++
			}
		}
		diffTempAlert()
	}
}

// OpenWeatherMapの気温とRaspberry Piの温度センサの気温に差がある場合、エラーメールを送信する
func diffTempAlert() {
	owmData := getOpenWeatherMapTemp()         // OpenWeatherMapから現在のデータを取得
	targetData := getNearTimeTempData(owmData) // OpenWeatherMapのデータの取得時刻に近い値を取得

	tempDiff := math.Abs(float64(owmData.Temp - targetData.Temp)) // 温度差の絶対値を計算する
	if tempDiff > 5.0 {                                           // 5℃以上気温差がある場合にアラートを出力する
		text := fmt.Sprint("Raspberry Pi   ", targetData, "\r\n") +
			fmt.Sprint("OpenWeatherMap ", owmData, "\r\n") +
			fmt.Sprint("温度差:", tempDiff, "\r\n")
		fmt.Print(text) // 標準出力に表示
		sendMail(text)  // メールで送信
	}
}

func getOpenWeatherMapTemp() TempData {
	getData, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Tokyo,jp&units=metric&appid=[appid]") // [appid]は実際の値に置換して使用
	defer getData.Body.Close()
	if err != nil {
		fmt.Println(err)
		return TempData{}
	}

	jsonData, err := ioutil.ReadAll(getData.Body)
	if err != nil {
		fmt.Println(err)
		return TempData{}
	}

	js, _ := simplejson.NewJson(jsonData)
	temp, _ := js.GetPath("main", "temp").Float64() // キーmain.tempの値をfloat64型として取得
	utc, _ := js.GetPath("dt").Int64()              // キーdtの値をint64型として取得
	date := time.Unix(utc, 0)                       // dtから取得した値をtime.Unix関数によりtime.Time型に変換

	return TempData{float32(temp), date}
}

// 引数に一番近い時間のデータをtempDataSliceから検索する
func getNearTimeTempData(t TempData) TempData {
	tempDataMutex.Lock()
	defer tempDataMutex.Unlock() // 変数へのアクセスが終わったらロック解除する

	timeSub := math.Abs(float64(t.EntryTime.Sub(tempDataSlice[0].EntryTime)))
	targetData := tempDataSlice[0]
	for _, t := range tempDataSlice {
		timeSubTemp := math.Abs(float64(t.EntryTime.Sub(t.EntryTime)))
		if timeSub > timeSubTemp {
			timeSub = timeSubTemp
			targetData = t
		}
	}
	return targetData
}

// 引数で渡された温度をメールで送信する
func sendMail(mailText string) {
	hostname := "mail.example.com"
	auth := smtp.PlainAuth(
		"",
		"foo@example.com", // ユーザ名
		"password",        // パスワード
		hostname)          // SMTPサーバのホスト名
	fromaddr := "foo@example.com"
	toaddr := "bar@example.com"
	err := smtp.SendMail(
		hostname+":25", // SMTPサーバのホスト名およびポート
		auth,
		fromaddr,         // FROMアドレス
		[]string{toaddr}, // TOアドレスのスライス
		[]byte("From: "+fromaddr+"\r\n"+
			"To: "+toaddr+"\r\n"+
			"Subject: 温度差エラー通知\r\n"+
			"\r\n"+
			mailText)) // メールヘッダおよび本文
	if err != nil {
		fmt.Println(err)
		return
	}
}
