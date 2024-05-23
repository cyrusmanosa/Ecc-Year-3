package main

import "fmt"

func main() {

	// ループで足し算した結果  を格納する変数
	n := 0

	// 1から10までを足し算するループ
	for i := 1; i <= 10; i++ {
		n += i
	}

	// 結果表示
	fmt.Println(n)
}
