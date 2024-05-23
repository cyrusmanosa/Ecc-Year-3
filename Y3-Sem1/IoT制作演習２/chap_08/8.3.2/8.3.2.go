package main

import "fmt"

func main() {
	n := 10

	// 外側の変数nを出力し、その後インクリメントする
	f := func() {
		fmt.Println(n)
		n++
	}

	f()            // 10を出力、nをインクリメント
	fmt.Println(n) // 11を出力

	n = 20
	f()            // 20を出力、nをインクリメント
	fmt.Println(n) // 21を出力
}
