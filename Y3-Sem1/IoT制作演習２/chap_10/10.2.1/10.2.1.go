package main

import "fmt"

// typeキーワードによるインタフェース型の宣言
type myIF interface {
	typeName() string
	add(n int) int
}

// myIFインタフェースを実装する型の宣言(数値型)
type myInt int

// myInt型のtypeNameメソッドの宣言
func (m myInt) typeName() string {
	return "myInt"
}

// myInt型のaddメソッドの宣言
func (m myInt) add(n int) int {
	return int(m) + n
}

// myInt型のsubメソッドの宣言
func (m myInt) sub(n int) int {
	return int(m) - n
}

// myIFインタフェースを実装する型の宣言(文字列型)
type myString string

// myString型のtypeNameメソッドの宣言
func (m myString) typeName() string {
	return "myString"
}

// myString型のaddメソッドの宣言
func (m myString) add(n int) int {
	return len(m) + n
}

func main() {
	// myInt型の値をmyIF型の変数に設定
	var n myInt = 1
	var i myIF = n
	fmt.Println(i.typeName()) // myInt型のtypeNameメソッドを呼び出す
	fmt.Println(i.add(2))     // myInt型のaddメソッドを呼び出す
	// i.sub(3)               // myIF型からはsubメソッドを呼び出せない

	// ポインタも設定可能
	i = &n
	fmt.Println(i.typeName()) // myInt型のtypeNameメソッドを呼び出す
	fmt.Println(i.add(2))     // myInt型のaddメソッドを呼び出す

	// myString型の値をmyIF型の変数に設定
	var s myString = "abc"
	i = s
	fmt.Println(i.typeName()) // myString型のtypeNameメソッドを呼び出す
	fmt.Println(i.add(2))     // myString型のaddメソッドを呼び出す
}
