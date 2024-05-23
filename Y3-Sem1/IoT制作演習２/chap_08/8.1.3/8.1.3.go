package main

import "fmt"

// 可変長引数のみの関数宣言
func funcVariadic1(slice ...string) {
	// 引数の中身を出力
	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
}

// 固定の引数と可変長引数を持つ関数の宣言
func funcVariadic2(n int, slice ...string) {
	// 引数の中身を出力
	fmt.Print(n, ": ")
	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
}

func main() {
	// 可変長引数を持つ関数の呼び出し
	funcVariadic1()
	funcVariadic1("a")
	funcVariadic1("a", "b")

	// 可変長引数を持つ関数の呼び出し（実引数にスライスを使用)
	slice := []string{"a", "b"}
	funcVariadic1(slice...)

	// 固定の引数と可変長引数を持つ関数の呼び出し
	funcVariadic2(10, "a", "b")
}
