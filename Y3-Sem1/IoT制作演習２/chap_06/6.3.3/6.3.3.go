package main

import "fmt"

func main() {

    // 計算結果を格納する変数
    n := 1

    // 1000を超えるまでnを2倍にし続ける
    for n <= 1000 {
        n *= 2
    }

    // 結果表示
    fmt.Println(n)

}
