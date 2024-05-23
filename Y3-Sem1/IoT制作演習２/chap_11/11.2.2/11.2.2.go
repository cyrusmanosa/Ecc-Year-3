package main

import "fmt"

func main() {
	// マップ型の変数宣言
	map1 := map[string]int{"a": 0, "b": 1}

	// キーが存在する場合
	n, ok := map1["a"]
	fmt.Println(n, ok)

	// キーが存在しない場合
	n, ok = map1["c"]
	fmt.Println(n, ok)
}
