package main

import "fmt"

// 戻り値を持つ関数(戻り値1つ)
func funcResult1() int {
	return 1
}

// 戻り値を持つ関数(戻り値2つ)
func funcResult2() (int, int) {
	return 1, 2
}

// 引数を持つ関数(引数2つ)
func funcResult3(x, y int) {
	fmt.Println(x + y)
}

func main() {
	// 戻り値を変数に設定
	a := funcResult1()
	fmt.Println(a)

	// 複数の戻り値を変数に設定
	a, b := funcResult2()
	fmt.Println(a, b)

	// 関数呼び出しの結果を実引数に使用
	funcResult3(funcResult2())
}
