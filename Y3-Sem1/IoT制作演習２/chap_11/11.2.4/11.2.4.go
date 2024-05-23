package main

import (
	"fmt"
	"time"
)

// ループ回数
const c = 100000

// マップにキーと要素を設定するループ処理
func funcMapLoop(m map[int]int) {
	for n := 0; n < c; n++ {
		m[n] = 1
	}
}

func main() {
	// 要素の数を指定しないmake関数
	t1 := time.Now()                     // 現在時刻を取得
	map1 := make(map[int]int)            // 要素の数を指定しない
	funcMapLoop(map1)                    // マップにキーと要素を設定
	fmt.Println("map1:", time.Since(t1)) // 処理時間を出力

	// 要素の数を指定するmake関数
	t2 := time.Now()                     // 現在時刻を取得
	map2 := make(map[int]int, c)         // 要素の数を指定
	funcMapLoop(map2)                    // マップにキーと要素を設定
	fmt.Println("map2:", time.Since(t2)) // 処理時間を出力
}
