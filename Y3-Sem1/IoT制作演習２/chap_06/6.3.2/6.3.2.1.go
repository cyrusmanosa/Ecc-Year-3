package main

import "fmt"

func main() {

	// continue文の例
	for i := 1; i <= 10; i++ {
		// iの値が偶数の場合は出力せず、次の繰り返し処理を実行
		if (i % 2) == 0 {
			continue
		}
		fmt.Print(i, ",")
	}
	fmt.Println()

	// break文の例
	for i := 1; i <= 10; i++ {
		// iの値が「5」のとき繰り返し処理を終了
		if i == 5 {
			break
		}
		fmt.Print(i, ",")
	}
	fmt.Println()
}
