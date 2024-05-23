package main

import "fmt"

func main() {
	s1 := "abc"
	s2 := "あいうえお"

	// 文字列「abc」と「あいうえお」の結合
	s3 := s1 + s2
	fmt.Println(s3)         // abcあいうえお

	// 文字列「abcあいうえお」のバイト長
	fmt.Println(len(s3))    // 18
}
