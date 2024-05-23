package main

import "fmt"

// 構造体型の宣言
type point struct {
	name string
	x, y int64
}

// 匿名フィールドを使用した構造体の宣言
type point3 struct {
	point // 匿名フィールド
	y, z int64
}

func main() {
	// 匿名フィールドを持つ構造体型の変数
	var p1 point3
	p1.name = "p1name" // 「p1.point.name = "p1name"」と同じ意味
	p1.x = 1           // 「p1.point.x = 1」と同じ意味
	p1.y = 2           // 「p1.point.y = 2」と別の意味
	p1.point.y = 20
	p1.z = 3

	fmt.Println("name:", p1.name, p1.point.name)
	fmt.Println("x:", p1.x, p1.point.x)
	fmt.Println("y:", p1.y, p1.point.y)
	fmt.Println("z:", p1.z)
}
