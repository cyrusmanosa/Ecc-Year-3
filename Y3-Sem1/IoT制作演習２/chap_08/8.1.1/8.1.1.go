package main

import "fmt"

// funcName1の関数の宣言
func funcName1() {
	fmt.Println("Hello")
}

func main() {
	// funcName1の呼び出し
	funcName1()

	// 変数funcName2にfuncName1の関数を設定
	funcName2 := funcName1

	// funcName2の呼び出し
	funcName2()
}
