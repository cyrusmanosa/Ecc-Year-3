package main

import "fmt"

func main() {
	s := "aあいう"

	// マルチバイト文字単位で繰り返す処理を実行する
	for i, r := range s {
		// バイト単位の位置、Unicodeコードポイント、文字列を表示
		fmt.Println(i, r, string(r))
	}
}
