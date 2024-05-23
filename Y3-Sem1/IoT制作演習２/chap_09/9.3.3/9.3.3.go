package main

import "fmt"

func main() {
	// ポインタ型の変数にnilを設定
	var x *int = nil

	fmt.Println(x)      // ポインタ自体の参照はエラーにならない
	fmt.Println(*x)     // この行でエラーが発生
}
