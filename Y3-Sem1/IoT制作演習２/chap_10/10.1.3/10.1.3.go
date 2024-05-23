package main

import "fmt"

// 名前付きの型を作成
type myInt int

// 値レシーバを持つメソッドの宣言
func (m myInt) addValue(n myInt) {
	m += n
}

// ポインタレシーバを持つメソッドの宣言
func (m *myInt) addPointer(n myInt) {
	*m += n
}

func main() {
	// 値レシーバを持つメソッドの呼び出し
	var v myInt = 1
	v.addValue(2)  // vの値は変更されない
	fmt.Println(v) // 1を出力

	// ポインタレシーバを持つメソッドの呼び出し
	var p *myInt = &v
	p.addPointer(2) // ポインタ先の値は変更される
	fmt.Println(*p) // 3を出力
}
