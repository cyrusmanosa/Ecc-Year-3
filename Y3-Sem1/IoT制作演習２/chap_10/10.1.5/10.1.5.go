package main

import "fmt"

// 名前付きの型を作成
type myInt int

// メソッドの宣言
func (m myInt) method(s string) {
	fmt.Println(s, m)
}

func main() {
	// メソッド値
	var v myInt = 1
	methodValue := v.method // 変数の型は「func(s string)」
	methodValue("メソッド値:")

	// 値のメソッド式
	methodExpr := myInt.method // 変数の型は「func(m myInt, s string)」
	methodExpr(2, "値のメソッド式:")

	// ポインタのメソッド式
	methodExprPtr := (*myInt).method // 変数の型は「func(m *myInt, s string)」
	v = 3
	methodExprPtr(&v, "ポインタのメソッド式:")
}
