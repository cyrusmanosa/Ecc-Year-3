package main

import "fmt"

// 実引数で呼び出される関数
func funcArg() string {
	fmt.Println("funcArg関数の実行")
	return "funcArgの戻り値"
}

// 実引数で関数呼び出し
func funcDefer1() {
	// 実引数のfuncArg関数は、defer文の実行時に処理
	defer fmt.Println(funcArg())
	fmt.Println("funcDefer1関数終了")
}

// 関数型変数を使用
func funcDefer2() {
	// defer文に登録する関数
	funcVar := func() { fmt.Println("defer文登録関数の実行") }

	// 登録する関数は、defer文の実行時に決まる
	defer funcVar()

	// defer文の実行後に変数を更新しても反映されない  
	funcVar = func() { fmt.Println("この関数は使用されない") }
}

func main() {
	funcDefer1()
	funcDefer2()
}
