package main

import "fmt"

func main() {

	// 計算結果を格納する変数
	n := 1

	// 1000を超えるまでnを2倍にし続ける
	for {
		n *= 2
		// 1000を超えたらfor文を脱出
		if 1000 <= n {
			break
		}
	}

	// 結果表示
	fmt.Println(n)
}
