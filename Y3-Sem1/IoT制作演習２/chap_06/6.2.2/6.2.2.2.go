package main

import "fmt"

func main() {
	x := 1
	y := 2

	// xの値で分岐
	switch x {
	case 1:
		fmt.Println("A")

		if y == 2 {
			// switch文のブロックから脱出
			break
		}

		// 下記の文は実行されない
		fmt.Println("B")
	}
}
