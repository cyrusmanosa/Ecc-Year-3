package main

import "fmt"

func main() {
	x := 1

	// xの値で分岐
	switch x {
	case 1:
		fmt.Println("A")

		// fallthrough文により、次のcaseである「case 2:」の中に移動
		fallthrough

	case 2:
		fmt.Println("B")
		// fallthrough文がない場合はswitch文のブロックを抜ける

	case 3:
		fmt.Println("C")
	}
}
