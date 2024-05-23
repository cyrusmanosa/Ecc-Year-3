package mypkg

func Add(n, m int) int {
	return n + m
}

// ループ回数
const c = 100000

// マップにキーと要素を設定するループ処理
func funcMapLoop(m map[int]int) {
	for n := 0; n < c; n++ {
		m[n] = 1
	}
}

func MapLoop1() {
	// 要素の数を指定しないmake関数
	map1 := make(map[int]int) // 要素の数を指定しない
	funcMapLoop(map1)         // マップにキーと要素を設定
}

func MapLoop2() {
	// 要素の数を指定するmake関数
	map2 := make(map[int]int, c) // 要素の数を指定
	funcMapLoop(map2)            // マップにキーと要素を設定
}
