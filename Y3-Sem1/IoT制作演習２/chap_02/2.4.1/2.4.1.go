package pkg1

import "fmt"

// init関数その1（プログラム開始時に実行）
func init() {
	fmt.Println("init1")  // 「init1」を出力
}

// init関数その2（プログラム開始時に実行）
func init() {
	fmt.Println("init2")    // 「init2」を出力
}
