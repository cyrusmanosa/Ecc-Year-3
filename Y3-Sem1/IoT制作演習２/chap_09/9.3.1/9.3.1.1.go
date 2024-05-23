package main

import "fmt"

func main() {

	// アドレスを取得する変数の宣言
	var v int32 = 100

	// アドレス演算子
	var pointer *int32 = &v // ポインタ型の変数に設定
	fmt.Println("ポインタの値:", pointer)

	// 間接演算子
	fmt.Println("間接演算子の結果:", *pointer)

	// 変数vの値を更新
	v = 200
	fmt.Println("変数vの更新:", *pointer) // 200を出力

	// ポインタ側から値を更新
	*pointer = 300
	fmt.Println("ポインタから値を更新:", v) // 300を出力
}
