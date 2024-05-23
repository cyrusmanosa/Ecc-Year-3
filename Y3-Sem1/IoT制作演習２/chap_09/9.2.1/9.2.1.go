package main

import "fmt"

// 構造体型の宣言
type point struct {
	name string
	x, y int64
}

func main() {
	// 各フィールドを指定して代入
	var p1 point
	p1.name = "point"
	p1.x = 1
	p1.y = 2
	fmt.Println("p1:", p1.name, p1.x, p1.y)

	// 構造体ごとの設定
	p2 := p1
	p1.name = "この変更はp2に反映されない"
	fmt.Println("p2:", p2.name, p2.x, p2.y)

	// 構造体リテラル（フィールド名の指定なし）
	p3 := point{"point", 1, 2}
	fmt.Println("p3:", p3.name, p3.x, p3.y)

	// 構造体リテラル（フィールド名の指定あり）
	p4 := point{name: "point", y: 2} // xは指定しないためゼロ値
	fmt.Println("p4:", p4.name, p4.x, p4.y)

	// 構造体リテラル（フィールドなし）
	p5 := point{}
	fmt.Println("p5:", p5.name, p5.x, p5.y)
}
