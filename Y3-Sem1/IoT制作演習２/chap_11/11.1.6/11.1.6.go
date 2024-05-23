package main

import "fmt"

type myStruct struct {
	s string
	x int
}

func main() {
	// 配列リテラル内の構造体リテラル
	s1 := [...]myStruct{myStruct{"A", 1}, myStruct{"B", 1}}
	s2 := [...]myStruct{{"A", 1}, {"B", 1}} // 上記の省略記法

	// 配列リテラル内の配列リテラル
	array1 := [...][2]int{[2]int{1, 2}, [2]int{3, 4}}
	array2 := [...][2]int{{1, 2}, {3, 4}} // 上記の省略記法

	// 配列リテラル内のスライスリテラル
	slice1 := [...][]int{[]int{1, 2}, []int{3, 4}}
	slice2 := [...][]int{{1, 2}, {3, 4}} // 上記の省略記法

	// 配列リテラル内のマップリテラル
	map1 := [...]map[int]int{map[int]int{1: 2}}
	map2 := [...]map[int]int{{1: 2}} // 上記の省略記法

	fmt.Println(s1, s2)
	fmt.Println(array1, array2)
	fmt.Println(slice1, slice2)
	fmt.Println(map1, map2)
}
