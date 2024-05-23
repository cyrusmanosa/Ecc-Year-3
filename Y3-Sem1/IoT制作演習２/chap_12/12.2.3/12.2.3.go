package main

import "fmt"

func funcPanic() {
	fmt.Println("[funcPanic start]")
	panic("[funcPanic panic]")     // パニック発生
	fmt.Println("[funcPanic end]") // パニック発生のため、実行されない
}

func funcRecover() {
	fmt.Println("[funcRecover start]")
	defer func() {
		// recover関数の呼び出しで通常の処理の流れに戻る
		r := recover()
		if r != nil {
			fmt.Println("[recover]", r)
		}
	}()

	funcPanic()
	fmt.Println("[funcRecover end]") // パニック発生中のため実行されない
}

func main() {
	fmt.Println("[main start]")
	funcRecover()
	fmt.Println("[main end]") // パニックから回復したため、実行される
}
