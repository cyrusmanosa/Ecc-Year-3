package main

import "fmt"

func main() {
	// マップ型の変数宣言
	map1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// マップのfor文:1回目
	for k, e := range map1 {
		fmt.Print(k, ":", e, ", ")
	}
	fmt.Println() // 改行

	// マップのfor文:2回目
	// ループする順序がランダムのため1回目と同じにならない
	for k, e := range map1 {
		fmt.Print(k, ":", e, ", ")
	}
	fmt.Println() // 改行
}
