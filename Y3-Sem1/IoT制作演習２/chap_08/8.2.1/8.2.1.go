package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func printIf(b bool) {
	// 関数型の変数を宣言
	var funcVar func(int, int) int

	// 下記のように、仮引数名を付けて書くことも可能
	// var funcVar func(x, y int) int

	// 条件分岐で設定する関数を変更する
	if b {
		funcVar = add
	} else {
		funcVar = mul
	}

	// 設定した関数の呼び出し
	fmt.Println(funcVar(3, 4))
}

func main() {
	printIf(true)   // 7を出力
	printIf(false)  // 12を出力
}
