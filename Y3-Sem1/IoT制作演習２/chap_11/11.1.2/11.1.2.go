package main

import "fmt"

// スライスの要素を出力する関数
func printSliceInt(s []int) {
	if s == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("len:", len(s), s[0], s[1])
	}
}

func main() {
	// int型スライスの変数宣言
	var slice1 []int // nilが設定される

	// スライスリテラルによる初期値の設定
	slice2 := []int{10, 20, 30} // 長さ3のint型スライス
	slice3 := []int{10, 20}     // 長さ2のint型スライス

	// すべて同じ型のため、引数に使用できる
	printSliceInt(slice1)
	printSliceInt(slice2)
	printSliceInt(slice3)
}
