package main

import (
	"fmt"
	"runtime"
	"time"
)

// ループ回数
const c = 10000000

// 時間のかかるループ処理
func funcGoLoop(ch chan<- int) {
	n := 0
	for i := 0; i < c; i++ {
		n += i
	}
	ch <- n // 処理が終了したことを知らせる
}

// 時間のかかる2つのゴルーチンを実行
func funcGo() {
	ch := make(chan int) // 待ち合わせ用のチャネル

	// 2つのゴルーチンを実行することで
	// 可能であればOSスレッドを2つ使用する
	go funcGoLoop(ch)
	go funcGoLoop(ch)

	// 2つのゴルーチンが終了するまで待つ
	<-ch
	<-ch
}

func main() {
	// 使用するOSスレッドの最大数を「1」に設定
	runtime.GOMAXPROCS(1)
	t1 := time.Now() // 現在の時刻を取得
	funcGo()
	fmt.Println("GOMAXPROCS=1:", time.Since(t1)) // 処理時間を出力

	// 使用するOSスレッドの最大数を「2」に設定
	runtime.GOMAXPROCS(2)
	t2 := time.Now() // 現在の時刻を取得
	funcGo()
	fmt.Println("GOMAXPROCS=2:", time.Since(t2)) // 処理時間を出力
}
