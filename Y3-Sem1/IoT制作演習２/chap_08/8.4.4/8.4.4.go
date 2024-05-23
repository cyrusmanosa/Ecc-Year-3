package main

import "fmt"

// defer文で戻り値を変更する関数
func funcDeferResult() (result int) {
	// defer文に関数リテラルの呼び出しを登録
	defer func() {
		// 名前付き戻り値を出力
		fmt.Println("更新前戻り値:", result)
		// 名前付き戻り値を更新
		result += 2
	}()

	// return文で戻り値を指定
	return 1
}

func main() {
	// 戻り値を出力
	fmt.Println("関数の戻り値:", funcDeferResult())
}
