package main

import "fmt"

type MyString string

// MyString型のStringメソッド
func (s MyString) String() string {
	return "***" + string(s) + "***"
}

func main() {
	var s MyString = "ABC"

	// Print関数は、MyString型のStringメソッドの戻り値を出力する
	fmt.Print(s)
}
