package main

import "fmt"

func main() {
	// マップ型変数の宣言（マップリテラル使用）
	map1 := map[string]int{"a": 1, "b": 2}

	// マップの要素の代入
	map1["a"], map1["c"] = 3, 4
	fmt.Println("len(map1):", len(map1))
	fmt.Println("map1:", map1["a"], map1["b"], map1["c"], map1["d"])

	// マップの代入
	map2 := map1
	map1["c"] = 5 // マップは参照型のため、map2の要素も更新する
	fmt.Println("map2:", map2["a"], map2["b"], map2["c"])

	// マップ型変数の宣言（初期値指定なし）
	var map3 map[string]int         // 初期値はnil
	fmt.Println("map3:", map3["a"]) // nilマップの要素は常にゼロ値
	map3["a"] = 1                   // nilマップの要素に代入するとエラー
}
