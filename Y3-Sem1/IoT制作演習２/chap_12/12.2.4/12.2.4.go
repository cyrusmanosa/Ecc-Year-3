package main

import "fmt"

func funcRuntimePanic() {
	var a []int       // 初期値はnil
	fmt.Println(a[0]) // a[0]でランタイムパニック発生
}

func main() {
	defer func() {
		// recover関数の呼び出しで通常の処理の流れに戻る
		r := recover()
		if r != nil {
			fmt.Println("[recover]", r)
		}
	}()

	funcRuntimePanic()
}
