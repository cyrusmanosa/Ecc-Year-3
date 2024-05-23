package main

import "fmt"

// 関数を戻り値とする関数
func funcResultFunc() func() {
	n := 10

	// 戻り値となる関数リテラル
	return func() {
		// 外側のローカル変数nを出力する
		fmt.Println(n)
	}
}

func main() {
	f := funcResultFunc() // 戻り値として関数を受け取る
	f()                   // 受け取った関数を実行
}
