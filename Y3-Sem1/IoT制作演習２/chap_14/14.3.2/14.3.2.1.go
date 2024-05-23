package main

import "fmt"

// 共有変数の宣言
var sharedVar string

// 共有変数に値を書き込む関数
func writeSharedVar(s string) {
	sharedVar = s
}

func readSharedVar() string {
	return sharedVar
}

func main() {
	// 共有変数に書き込む変数の宣言
	s1 := "AB"
	s2 := "C"
	fmt.Println("s1:", s1, len(s1))
	fmt.Println("s2:", s2, len(s2))

	// 書き込み用ゴルーチンの起動
	go func() {
		// このゴルーチンでは、s1とs2を交互に共有変数に書き込む
		for {
			writeSharedVar(s1)
			writeSharedVar(s2)
		}
	}()

	for {
		// 特定のタイミングの共有変数の値を取り出す
		s3 := readSharedVar()

		// s1と長さが同じ、かつ、先頭の値がs2と同じ
		if len(s1) == len(s3) && s2[0] == s3[0] {
			fmt.Println("s3:", s3, len(s3)) // 不正な文字列
			break
		}
	}
}
