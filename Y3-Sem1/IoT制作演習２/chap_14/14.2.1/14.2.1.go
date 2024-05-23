package main

import "fmt"

// 送信側のゴルーチン用関数
func funcSend(ch chan<- int) {
	fmt.Println("[送信 10]")
	ch <- 10 // チャネルによる送信
}

func main() {
	ch := make(chan int) // バッファなしチャネル作成
	go funcSend(ch)      // ゴルーチン作成
	n := <-ch            // チャネルによる受信
	fmt.Print("[受信 ", n, "]")
}
