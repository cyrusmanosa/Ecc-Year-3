package main

import "fmt"

func funcPanic1() {
	fmt.Println("[funcPanic1 start]")
	panic("[funcPanic1 panic]")     // パニック発生
	fmt.Println("[funcPanic1 end]") // パニック発生のため、実行されない
}

func funcPanic2() {
	fmt.Println("[funcPanic2 start]")
	funcPanic1()                    // この関数呼び出し内でパニック発生
	fmt.Println("[funcPanic2 end]") // パニック発生のため、実行されない
}

func main() {
	fmt.Println("[main start]")
	funcPanic2()              // この関数呼び出し内でパニック発生
	fmt.Println("[main end]") // パニック発生のため、実行されない
}
