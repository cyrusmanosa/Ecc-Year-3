package main

import "fmt"

// 共有変数（グローバル変数）
var n1 = 0

func main() {
	// 共有変数（ローカル変数）
	n2 := 0

	// ゴルーチン待ち合わせ用のチャネル
	ch := make(chan struct{})

	// どちらのゴルーチンでも同じメモリアドレスを持つ
	go func() {
		n1, n2 = 1, 2
		fmt.Println("go:  ", n1, &n1, n2, &n2)
		close(ch)
	}()

	<-ch // go文のゴルーチンが終了するまで待つ
	fmt.Println("main:", n1, &n1, n2, &n2)
}
