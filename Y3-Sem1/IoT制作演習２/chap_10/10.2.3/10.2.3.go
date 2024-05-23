package main

import "fmt"

func main() {
	// 数値1を異なる型(int型,byte型)に変換した後に設定を設定
	var i1 interface{} = int(1)
	var i2 interface{} = byte(1)

	// 値が一致していても、型が異なる場合は不一致と判定
	fmt.Println("i1 == i2:", i1 == i2) // false

	// 異なる型のnilを設定
	var i3 interface{} = nil         // 型指定なしのnil
	var i4 interface{} = (*int)(nil) // *int型のnil

	// 両方ともnilでも、型が異なる場合は不一致と判定
	fmt.Println("i3 == i4:", i3 == i4) // false

	// 型指定なしのnilとの演算
	fmt.Println("i3 == nil:", i3 == nil) // true
	fmt.Println("i4 == nil:", i4 == nil) // false

	// *myInt型のnilとの演算
	fmt.Println("i3 == (*int)(nil):", i3 == (*int)(nil)) // false
	fmt.Println("i4 == (*int)(nil):", i4 == (*int)(nil)) // true
}
