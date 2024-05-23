package main

import "fmt"

func main() {

    // ローカル変数xの宣言(関数mainのブロックがスコープ)
    x := 1

    if x == 1 { // (1)

        // ローカル変数yの宣言(if文(1)のブロックがスコープ)
        y := 2
        fmt.Println(x, y)

    } // yを参照できる範囲はここまで

} // xを参照できる範囲はここまで
