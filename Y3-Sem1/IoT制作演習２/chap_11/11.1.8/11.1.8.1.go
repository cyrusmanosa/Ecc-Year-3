package main

import "fmt"

func main() {
	// 文字列にインデックス、スライス式を使用
	s1 := "aあ"
	fmt.Println(s1[0])               // 「a」のバイト単位の値
	fmt.Println(s1[1], s1[2], s1[3]) // 「あ」のバイト単位の値
	fmt.Println(s1[1:4])             // スライス式の結果は文字列「あ」
}
