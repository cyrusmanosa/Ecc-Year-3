package main

import "fmt"

//値を3回送信した後、チャネルをクローズする
func funcSendFor(ch chan<- int) {
	for n := 0; n < 3; n++ {
		ch <- n
	}
	close(ch) // チャネルクローズ
}

func main() {
	ch := make(chan int) // チャネル作成
	go funcSendFor(ch)   // ゴルーチン作成

	// クローズするまでループする
	for n := range ch {
		fmt.Println("受信:", n)
	}
}
