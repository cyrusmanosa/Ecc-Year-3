package main

import "fmt"

func main() {
	fmt.Printf("%v %v", 1) // 引数の数が足りない

	return
	fmt.Println("ABC") // returnが前にあるため、到達しない
}
