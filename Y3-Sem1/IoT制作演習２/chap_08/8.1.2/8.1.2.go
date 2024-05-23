package main

import "fmt"

// 引数を持つ関数の宣言
func funcParam1(n int) {
	fmt.Println(n)
}

// 引数を持つ関数の宣言（引数2つ)
func funcParam2(n1, n2 int) {
	fmt.Println(n1 + n2)
}

func main() {
	// 関数の呼び出し
	funcParam1(1)       // 1が実引数
	funcParam2(1, 2)    // 1と2が実引数
}
