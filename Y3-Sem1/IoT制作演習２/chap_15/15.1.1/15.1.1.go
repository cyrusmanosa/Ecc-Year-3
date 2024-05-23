package main

import "fmt"

func main() {
	// %の部分を2番目以降の引数で置き換える
	fmt.Printf("[%v]\n", "ABC")       // デフォルトの形式
	fmt.Printf("[%#v]\n", "ABC")      // ダブルクォート付き
	fmt.Printf("[%+d]\n", 10)         // 符号付き
	fmt.Printf("[%4d]\n", 10)         // 左パディング（空白埋め）
	fmt.Printf("[%04d]\n", 10)        // 左パディング（0埋め）
	fmt.Printf("[%6.2f]\n", 1234.567) // 小数点以下の精度指定
}
