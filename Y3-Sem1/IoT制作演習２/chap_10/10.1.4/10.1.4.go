package main

import "fmt"

// 名前付きの型を作成
type myInt int

// 値レシーバを持つメソッドの宣言
func (m myInt) methodValue() {
	fmt.Println("値レシーバ:", m)
}

// ポインタレシーバを持つメソッドの宣言
func (m *myInt) methodPointer() {
	if m == nil {
		fmt.Println("ポインタレシーバ: nil")
	} else {
		fmt.Println("ポインタレシーバ:", *m)
	}
}

func main() {
	// ポインタ型の変数から、値レシーバのメソッドの呼び出し
	var v myInt = 1
	var p *myInt = &v
	p.methodValue() // (*p).methodValue()と同じ意味

	// ポインタ型ではない変数から、ポインタレシーバのメソッドの呼び出し
	v.methodPointer() // (&v).methodPointer()と同じ意味

	// nilの場合のメソッド呼び出し
	p = nil
	p.methodPointer() // nilポインタから、ポインタレシーバのメソッドの呼び出し
	p.methodValue()   // nilポインタから、値レシーバ  のメソッドの呼び出し(エラー)
}
