package main

import "fmt"

func main() {

	// goto文でラベルにジャンプする
	goto label

	// 実行されない文
	fmt.Println("この文は表示されない")

label:
	fmt.Println("goto文による移動")
}
