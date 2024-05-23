package main

import "fmt"

func main() {
	// 関数リテラルを使用した関数の作成
	f := func(i, j int) int {
		return i + j
	}

	// 作成した関数の呼び出し
	n := f(1, 2)
	fmt.Println(n)
}
