package main

import (
	"fmt"
	"time"
)

// 送信側の関数
func funcSend(ch chan<- int) {
	// 1ミリ秒ごとに値を3回送信する
	for n := 0; n < 3; n++ {
		time.Sleep(time.Millisecond) // 1ミリ秒待つ
		fmt.Print("[送信 ", n, "]")
		ch <- n // チャネルによる送信
	}
}

// 受信側の関数
func funcRecv(ch <-chan int) {
	// 10ミリ秒ごとに値を3回受信する
	for n := 0; n < 3; n++ {
		time.Sleep(10 * time.Millisecond) // 10ミリ秒待つ
		fmt.Print("[受信 ", <-ch, "]")      // チャネルによる受信
	}
}

func main() {
	// バッファ付きチャネル（容量3）の場合
	fmt.Print("バッファ付き（容量3）: ")
	ch1 := make(chan int, 3) // バッファ付きチャネル作成
	go funcSend(ch1)
	funcRecv(ch1)

	// バッファ付きチャネル（容量1）の場合
	fmt.Print("\nバッファ付き（容量1）: ")
	ch2 := make(chan int, 1) // バッファ付きチャネル作成
	go funcSend(ch2)
	funcRecv(ch2)

	// バッファなしチャネルの場合
	fmt.Print("\nバッファなし（容量0）: ")
	ch3 := make(chan int) // バッファなしチャネル作成
	go funcSend(ch3)
	funcRecv(ch3)
}
