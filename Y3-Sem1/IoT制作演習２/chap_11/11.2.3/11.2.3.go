package main

import "fmt"

func main() {
	// マップ型の変数宣言
	map1 := map[string]int{"a": 0, "b": 1}

	// マップからキー"a"を削除
	delete(map1, "a")

	// マップにキー"a"が存在しないため、変数okはfalseとなる
	n, ok := map1["a"]
	fmt.Println(n, ok)
}
