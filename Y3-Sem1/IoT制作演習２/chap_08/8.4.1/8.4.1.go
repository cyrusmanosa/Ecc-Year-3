package main

import "fmt"

func main() {
	fmt.Println("main関数開始")

	// defer文で登録した関数呼び出しは、関数終了時に実行
	defer fmt.Println("defer文の関数呼び出し実行")

	fmt.Println("main関数終了")
}
