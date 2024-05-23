package main

import (
	"encoding/json"
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"net/http"
	"os"
	"os/signal"
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

// main関数では温度取得の開始およびWebサーバの起動を行う
func main() {
	wg.Add(1)
	go getTemp() // 温度データ取得のゴルーチンを開始

	wg.Add(1)
	go setLed() // LED点滅のゴルーチンを開始

	wg.Wait() // 温度センサおよびLEDの設定を待つ

	// Webサーバの設定と起動
	http.HandleFunc("/json", getJsonData)
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
	intervalTime := 10
	for {
		// 以下のfor文により10秒待機を実現する … ⑤
		i := 0
		for i < intervalTime { // 1 * intervalTime秒待つ
			select {
			case <-tick: // 1秒ごとにインクリメントする
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
