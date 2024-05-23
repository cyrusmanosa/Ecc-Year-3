package main

import "fmt"

// "[trueFunc]"を出力してtrueを返す関数
func trueFunc() bool {
	fmt.Print("[trueFunc]")
	return true
}

// "[falseFunc]"を出力してfalseを返す関数
func falseFunc() bool {
	fmt.Print("[falseFunc]")
	return false
}

func main() {
	// 左側(trueFunc)のみ評価される。右側は評価されない。
	fmt.Println(trueFunc() || falseFunc())

	// 左側(falseFunc)のみ評価される。右側は評価されない。
	fmt.Println(falseFunc() && trueFunc())

	// 両方とも評価される。
	fmt.Println(falseFunc() || falseFunc())
	fmt.Println(trueFunc() && trueFunc())
}
