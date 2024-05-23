package main

import "fmt"

type myStruct struct {
	field int
}

func main() {
	// 構造体リテラルにアドレス演算子を使用
	fmt.Println("&myStruct{1}:", &myStruct{1})

	// 構造体変数のフィールドにアドレス演算子を使用
	s := myStruct{1}
	fmt.Println("&s.field:", &s.field)

	// 配列アクセスにアドレス演算子を使用
	array := [10]int{}
	fmt.Println("&array[0]:", &array[0])

	// スライスのアクセスにアドレス演算子を使用
	slice := []int{0}
	fmt.Println("&slice[0]:", &slice[0])
}
