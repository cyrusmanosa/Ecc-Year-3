package main

import (
	"encoding/json"
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"html/template"
	"net/http"
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

	// Webサーバの設定と起動
	http.HandleFunc("/json", getJsonData)
	http.HandleFunc("/setting_interval", settingInterval) // /setting_intervalを追加
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

		// 最新の100件を残す … ④
		dataLimit := 100
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
