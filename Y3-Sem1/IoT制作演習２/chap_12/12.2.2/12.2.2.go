package main

import "fmt"

func funcPanicDefer1() {
	defer fmt.Println("[defer funcPanicDefer1]") // パニック発生中でも実行
	fmt.Println("[funcPanicDefer1 panic start]")
	panic("[funcPanicDefer1 panic]") // パニック発生
}

func funcPanicDefer2() {
	defer fmt.Println("[defer funcPanicDefer2]") // パニック発生中でも実行
	funcPanicDefer1()
}

func main() {
	defer fmt.Println("[defer main]") // パニック発生中でも実行
	funcPanicDefer2()
}
