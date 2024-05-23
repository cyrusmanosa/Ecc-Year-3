package main

import "fmt"

func main() {
	// 文字列を使用したappend関数
	b1 := make([]byte, 0, 5)
	b1 = append(b1, "abcde"...)
	fmt.Println(string(b1))

	// 文字列を使用したcopy関数
	copy(b1, "123")
	fmt.Println(string(b1))

	// 文字列連結の繰り返し(append関数)
	const size = 30
	const c = "a"
	b2 := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		b2 = append(b2, c...)
	}
	fmt.Println(string(b2))

	// 文字列連結の繰り返し(文字列の「+」演算子)
	s1 := ""
	for i := 0; i < size; i++ {
		s1 += c
	}
	fmt.Println(s1)
}
