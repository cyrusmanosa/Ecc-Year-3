package main

import "fmt"

// インタフェース型の宣言
type embedIF interface {
	method1()
}

// インタフェース型の宣言
type myIF interface {
	embedIF // 埋め込んだ  インタフェースの埋め込み
	method2()
}

// myIFインタフェースを実装する型の宣言
type myInt int

func (i myInt) method1() {
	fmt.Println("method1:", i)
}

func (i myInt) method2() {
	fmt.Println("method2:", i)
}

func main() {
	// インタフェース型の変数
	var i myIF = myInt(1)
	i.method1() // 埋め込んだ  embedIFのメソッド
	i.method2()
}
