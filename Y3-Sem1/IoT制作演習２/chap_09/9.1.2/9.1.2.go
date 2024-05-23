package main

type myInt1 int          // myInt1型の基底型はint型
type myInt2 myInt1       // myInt2型の基底型はint型
type myArray1 [10]int    // myArray1型の基底型は[10]int型
type myArray2 [10]myInt1 // myArray2型の基底型は[10]myInt1型

type namedType1 [10]int // 基底型は[10]int
type namedType2 [10]int // 基底型は[10]int

func main() {

	var unnamed [10]int     // 名前のない型の変数
	var named1 namedType1   // 名前付きの型の変数その1
	var named2 namedType2   // 名前付きの型の変数その2

	named1 = unnamed        // 名前のない型から名前付きの型には代入可能
	unnamed = named1        // 名前付きの型から名前のない型には代入可能
	//named1 = named2       // エラー。基底型が同じでも、異なる名前付きの型には代入できない

	named1 = namedType1(named2) // 基底型が同じ場合、型変換が可能


	var p1 *namedType1     // ポインタ型変数その1
	var p2 *namedType2     // ポインタ型変数その2
	p1 = (*namedType1)(p2) // 型変換が可能

	p2 = (*namedType2)(p1)
}
