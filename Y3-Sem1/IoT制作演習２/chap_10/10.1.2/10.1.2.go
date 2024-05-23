package main

import "fmt"

// 名前付きの型を作成
type myInt int
type myStruct struct {
	x, y int
}

// myInt型(数値型)のメソッドの宣言
func (m myInt) add(n int) myInt {
	return m + myInt(n)
}

// myStruct型(構造体型)のメソッドの宣言
func (m myStruct) add(n int) myStruct {
	return myStruct{m.x + n, m.y + n}
}

func main() {
	// myInt型のメソッド呼び出し
	var v myInt = 1
	fmt.Println(v.add(2))

	// myStruct型のメソッド呼び出し
	var s = myStruct{4, 5}
	fmt.Println(s.add(6))

	// 型の宣言ではメソッドを引き継がない。以下のコードを実行するとエラーになる
	// type myStruct2 myStruct
	// s2 := myStruct2{4, 5}
	// fmt.Println(s2.add(6))
}
