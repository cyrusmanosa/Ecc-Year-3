package main

import "fmt"

func main() {
	x := 1          // 「var x = 1」と同じ変数の宣言
	y, z := 2, 3    // 「var y, z = 2, 3」と同じ変数の宣言

	// 下記のように、型を指定する書き方はできない
	// x int := 1

	fmt.Println(x, y, z)
}
