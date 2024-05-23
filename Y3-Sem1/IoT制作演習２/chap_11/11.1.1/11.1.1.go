package main

import "fmt"

// 長さ3の配列を出力する関数
func printArrayInt(a [3]int) {
	fmt.Println(a[0], a[1], a[2])
}

func main() {
	// 初期値を指定しない変数宣言
	var array1 [3]int // すべてゼロ値が設定される

	// 初期値に配列リテラルを指定した変数宣言
	array2 := [3]int{20, 21}       // 長さを指定する配列（3番目の要素は0）
	array3 := [...]int{30, 31, 32} // 長さを指定しない配列
	array4 := [3]int{1: 41}        // インデックスを指定した配列

	printArrayInt(array1)
	printArrayInt(array2)
	printArrayInt(array3)
	printArrayInt(array4)
}
