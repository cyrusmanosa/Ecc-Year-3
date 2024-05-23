package main

import (
	"fmt"
	"strconv"
)

func main() {
	// iに10、errにnilが設定される
	i, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("[エラー出力]", err)
		return
	}
	fmt.Println(i + 1) // 11を出力

	// iに0、errにエラー値が設定される
	i, err = strconv.Atoi("あ")
	if err != nil {
		// 数値に変換できないのでエラー出力
		fmt.Println("[エラー出力]", err)
		return
	}
	fmt.Println(i + 1) // エラーのため、この行まで到達しない
}
