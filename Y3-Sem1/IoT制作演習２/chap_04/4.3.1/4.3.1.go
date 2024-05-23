package main

import "fmt"

// グローバル変数nの宣言
var n = 1

// main関数の宣言
func main() {
	fmt.Println(n)  // 1を出力
	f()             // nを2に更新
	fmt.Println(n)  // 2を出力
}

// f関数の宣言（nを2に更新）
func f() {
	n = 2
}
