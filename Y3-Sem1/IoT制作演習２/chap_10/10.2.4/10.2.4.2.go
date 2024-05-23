package main

import "fmt"

func typeAssertionCheck(i interface{}) {
	// int型の場合のみ、変数okにtrueを設定
	n, ok := i.(int)
	if ok {
		fmt.Println("int:", n+1)
	}

	// string型の場合のみ、変数okにtrueを設定
	s, ok := i.(string)
	if ok {
		fmt.Println("string: " + s)
	}
}

func main() {
	typeAssertionCheck(1)   // 「int: 2」を出力
	typeAssertionCheck("a") // 「string: a」を出力
}
