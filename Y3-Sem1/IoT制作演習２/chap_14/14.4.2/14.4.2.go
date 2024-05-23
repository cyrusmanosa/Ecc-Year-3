package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 使用するOSスレッドを1つにする
	runtime.GOMAXPROCS(1)

	// ゴルーチン実行
	go func() {
		fmt.Println("go")

		// 無限ループ。mainゴルーチンに切り替わらない
		for {
			// Gosched関数を呼び出すと、ゴルーチンが切り替わる
			// runtime.Gosched()
		}
	}()

	// mainゴルーチンの処理
	time.Sleep(time.Millisecond) // 1ミリ秒待つ
	fmt.Println("main")          // 実行されない
}
