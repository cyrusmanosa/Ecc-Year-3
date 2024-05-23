package main

import "fmt"

func main() {
	x := 1

	switch y := x + 1; y { // 変数yの有効範囲はここから  
	case 1:
		z := x + y // 変数zの有効範囲はこのcase内のみ
		fmt.Println("case 1:", x, y, z)

	case 2:
		// ここでは変数zを使用できない
		fmt.Println("case 2:", x, y)

	} // 変数yの有効範囲はここまで
}
