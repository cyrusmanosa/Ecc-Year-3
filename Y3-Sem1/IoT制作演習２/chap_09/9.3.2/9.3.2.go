package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	// new組込み関数の使用(数値型)
	n := new(int) // nはint型のポインタ型
	fmt.Println(n, *n)

	// new組込み関数の使用(構造体型)
	p1 := new(point) // p1はpoint型のポインタ型
	fmt.Println(*p1)

	// 複合リテラルのアドレス取得（new組込み関数と同じ意味）
	p2 := &point{} // p2はpoint型のポインタ型
	fmt.Println(*p2)
}
