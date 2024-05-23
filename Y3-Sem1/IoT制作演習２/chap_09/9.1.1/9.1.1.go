package main

import "fmt"

// int32型を使用した型の宣言
type myInt32 int32

// 丸カッコによる型の宣言
type (
	myInt    int
	myString string
)

func main() {
	// 型の宣言で作成したmyInt32型を使用
	var n1 myInt32 = 1
	var n2 myInt32 = 2
	fmt.Println(n1 + n2)

	// 異なる型の間の演算はできない
	// var i1 int32 = 1
	// fmt.Println(n1 + i1)
}
