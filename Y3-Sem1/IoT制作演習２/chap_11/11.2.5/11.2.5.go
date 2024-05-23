package main

import "fmt"

type myStruct struct {
	s string
	x int
}

func main() {
	// マップリテラル内の構造体リテラル
	// キーの構造体に省略記法を使用
	map1 := map[myStruct]int{
		{"A", 1}: 10, // 「myStruct{"A", 1}: 10」の省略記法
		{"B", 2}: 20, // 「myStruct{"B", 2}: 20」の省略記法
	}
	// 要素の構造体に省略記法を使用
	map2 := map[int]myStruct{
		10: {"A", 1}, // 「10: myStruct{"A", 1}」の省略記法
		20: {"B", 2}, // 「20: myStruct{"B", 2}」の省略記法
	}

	fmt.Println(map1[myStruct{"A", 1}], map1[myStruct{"B", 2}])
	fmt.Println(map2[10], map2[20])
}
