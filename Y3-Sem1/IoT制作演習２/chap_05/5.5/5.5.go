package main

import (
	"fmt"
	"unsafe"
)

// 定数式の例①(リテラル、定数同士の演算、型変換、組込み関数)
const (
	intLit = 123                            // 数値リテラル
	strLit = "ABC"                          // 文字列リテラル
	intConst = -intLit                      // 定数同士の演算(整数)
	strConst = strLit + "DEF"               // 定数同士の演算(文字列)
	floatConst = float64(intLit + 456)      // 型変換
	complexConst = complex(1.1, 2.2)        // complex組込み関数
	realConst = real(complexConst)          // real組込み関数
	imagConst = imag(complexConst)          // imag組込み関数
)

// len組込み関数、cap組込み関数で使用する配列
var a = [3]int{1, 2, 3}
// unsafeの組込み関数で使用する構造体変数
var s = struct{ x int64; y float64 }{}

// 定数式の例②(変数を引数にできる定数式)
const (
	lenConst = len(a)                       // len組込み関数
	capConst = cap(a)                       // cap組込み関数
	alignofConst = unsafe.Alignof(s)        // unsafe.Alignof組込み関数
	offsetofConst = unsafe.Offsetof(s.y)    // unsafe.Offsetof組込み関数
	sizeofConst = unsafe.Sizeof(s)          // unsafe.Sizeof組込み関数
)

func main() {
	// 定数の出力
	fmt.Println(intConst, strConst, floatConst)
	fmt.Println(complexConst, realConst, complexConst)
	fmt.Println(lenConst, capConst)
	fmt.Println(alignofConst, offsetofConst, sizeofConst)
}
