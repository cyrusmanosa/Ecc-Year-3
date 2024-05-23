package main

import "fmt"

func main() {
	// 配列のfor文
	array := []string{"a", "b", "c"}
	for i, e := range array {
		fmt.Print(i, ":", e, ",") // 配列のインデックスと要素を出力
	}
	fmt.Println()

	// スライスのfor文
	slice := []string{"d", "e", "f"}
	for i, e := range slice {
		fmt.Print(i, ":", e, ",") // スライスのインデックスと要素を出力
	}
	fmt.Println()

	// 変数が1つの場合はインデックスのみ
	for i := range array {
		fmt.Print(i, ",") // 配列のインデックスを出力
	}
}
