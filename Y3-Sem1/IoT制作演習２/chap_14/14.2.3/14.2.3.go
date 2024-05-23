package main

import "fmt"

//値を3回送信した後、チャネルをクローズする
func funcSendClose(ch chan<- int) {
	for n := 0; n < 3; n++ {
		ch <- n
	}

	close(ch) // チャネルクローズ
}

func main() {
	ch := make(chan int) // チャネル作成
	go funcSendClose(ch) // ゴルーチン作成

	for {
		n, ok := <-ch // 受信処理とクローズの確認

		// チャネルがクローズ済みの場合、ループを終了する
		if !ok {
			break
		}

		fmt.Println("受信:", n)
	}

	// クローズ済みの場合、ゼロ値を取得する
	fmt.Println("クローズ済:", <-ch)
}
