package main

import "fmt"

// 構造体型の宣言
type point2d struct {
	x, y int
}

// xとyの乗算
func (p2d point2d) mulXY() int {
	return p2d.x * p2d.y
}

// xとyの乗算(mulXYメソッドと同じ処理)
func (p2d point2d) mul() int {
	return p2d.x * p2d.y
}

// 匿名フィールドを持つ構造体型の宣言
type point3d struct {
	point2d // 匿名フィールド（mulXYメソッドとmulメソッドを引き継ぐ）
	z       int
}

// x、y、zの乗算（point3d特有のメソッド）
func (p3d point3d) mul() int {
	return p3d.x * p3d.y * p3d.z
}

func main() {
	p3d := point3d{point2d{2, 3}, 4}
	fmt.Println(p3d.mulXY())       // point2d型のmulXYメソッドを呼び出し
	fmt.Println(p3d.mul())         // point3d型のmulメソッドを呼び出し
	fmt.Println(p3d.point2d.mul()) // point2d型のmulメソッドを呼び出し
}
