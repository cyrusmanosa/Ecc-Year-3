package main

import "fmt"

// 名前付き戻り値の関数（値を指定しないreturn文)
func funcNamedResult1() (i int) {
	i = 10
	return // 10が戻り値
}

// 名前付き戻り値の関数（return文で値を指定する)
func funcNamedResult2() (i int) {
	i = 10
	return 20 // 20が戻り値(return文の値を優先するため)
}

// 名前付き戻り値の関数（複数の戻り値を持つ）
func funcNamedResult3() (i int, j int) {
	i, j = 10, 20
	return // 10, 20が戻り値
}

// 名前付き戻り値の関数（複数の戻り値を持つ。型をまとめて記述）
func funcNamedResult4() (i, j int) {
	i, j = 10, 20
	return // 10, 20が戻り値
}

// 型のゼロ値が戻り値となるケース
func funcNamedResult5() (i int) {
	return // 0が戻り値
}

func main() {
	fmt.Println(funcNamedResult1())
	fmt.Println(funcNamedResult2())
	fmt.Println(funcNamedResult3())
	fmt.Println(funcNamedResult4())
	fmt.Println(funcNamedResult5())
}
