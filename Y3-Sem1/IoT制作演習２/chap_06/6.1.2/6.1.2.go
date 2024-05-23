package main

import "fmt"

func main() {

    if x := 3; x == 1 { // xを宣言

        // xが「1」の場合に出力
        fmt.Println("x == 1")

    } else if x == 2 {

        // xが「2」の場合に出力
        fmt.Println("x == 2")

    } else {

        // すべてが異なる場合に出力
        fmt.Println(x)

    } // xの有効範囲はここまで
}
