package main

import "fmt"

func main() {

	// 整数型の比較演算
	fmt.Println(1 == 1, 1 != 2)
	fmt.Println(1 <= 2, 1 < 2)
	fmt.Println(2 >= 1, 2 > 1)

	// 文字列型の比較演算
	fmt.Println("ABC" == "ABC", "A" < "B", "ABC" < "ACB")
}
