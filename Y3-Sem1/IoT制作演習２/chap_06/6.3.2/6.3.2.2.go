package main

import "fmt"

func main() {

	// ラベル指定のcontinue文とbreak文の例
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			fmt.Println(i, j)

			if (i == 0) && (j == 1) {
				fmt.Println("continue")
				continue OuterLoop; // 外側のfor文の次の処理を実行
			}

			if (i == 1) && (j == 1) {
				fmt.Println("break")
				break OuterLoop; // 外側のfor文を終了する
			}
		}
	}
}
