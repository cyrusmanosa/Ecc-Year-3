package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	// Webサーバの設定と起動
	http.HandleFunc("/data", getData)
	http.ListenAndServe(":4000", nil)
}

// 温度データを取得し、表示
func getData(w http.ResponseWriter, r *http.Request) {
	bus := embd.NewI2CBus(1)
	bus.WriteByteToReg(0x48, 0x03, 0xA0) // ワンショットモードに設定
	time.Sleep(time.Second)              // 温度データの取得完了を待つ

	data16, _ := bus.ReadWordFromReg(0x48, 0x00) // 温度データを取得
	temp := float32(data16) / 128.0              // セ氏温度に変換

	fmt.Fprint(w, " 温度:", temp, "℃")
}
