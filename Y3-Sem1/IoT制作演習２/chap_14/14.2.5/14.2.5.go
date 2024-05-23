package main

import "fmt"

// ch1の送信処理とch2の受信処理を行うselect文（デフォルトなし）
func funcSelect(ch1 chan<- string, ch2 <-chan string) {
	select {

	// ch1送信処理
	case ch1 <- "a":
		fmt.Println("[select ch1送信: a]")

	// ch2受信処理
	case n := <-ch2:
		fmt.Println("[select ch2受信: " + n + "]")
	}
}

// ch1の送信処理とch2の受信処理を行うselect文（デフォルトあり）
func funcSelectDefault(ch1 chan<- string, ch2 <-chan string) {
	select {

	// ch1送信処理
	case ch1 <- "a":
		fmt.Println("[select ch1送信: a]")

	// ch2受信処理
	case n := <-ch2:
		fmt.Println("[select ch2受信: " + n + "]")

	// ch1の送信もch2の受信も実行できないケース
	default:
		fmt.Println("[select default]")
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { <-ch1 }() // ch1受信処理
	funcSelect(ch1, ch2)  // ch1送信処理を選択

	go func() { ch2 <- "b" }() // ch2送信処理
	funcSelect(ch1, ch2)       // ch2受信処理を選択

	// 送受信がないため、defaultを選択
	funcSelectDefault(ch1, ch2)
}
