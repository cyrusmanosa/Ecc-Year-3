package main

import "fmt"

func typeAssertion(i interface{}) {
	// int型に変換。iがint型以外の場合はエラー
	n := i.(int)
	fmt.Println(n + 1)
}

func main() {
	typeAssertion(1)
	typeAssertion("a") // 変換時に実行時エラー
}
