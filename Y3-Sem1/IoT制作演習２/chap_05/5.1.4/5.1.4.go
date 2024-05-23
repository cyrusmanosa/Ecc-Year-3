package main

import "fmt"

func main() {
	i := 1

	// インクリメント文（変数iに1を加算）
	i++
	fmt.Println(i)

	// デクリメント文（変数iに1を減算）
	i--
	fmt.Println(i)


	// 下記のような書き方はできない
	// a[i++]   // 配列のアクセス直後のインクリメント
	// b = i++  // 代入直後のインクリメント
}
