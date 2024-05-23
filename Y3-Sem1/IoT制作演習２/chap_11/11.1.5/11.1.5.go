package main

import "fmt"

func main() {

	// フルスライス式
	sliceA1 := []int{1, 2, 3, 4, 5}
	sliceA2 := sliceA1[0:2:2] // sliceA2の容量は2

	// sliceA2  の容量に空きがないため、
	// 要素追加時は新規にスライスを作成する
	// 新規にスライスを作成したため、sliceA1には影響しない
	sliceA3 := append(sliceA2, 6)
	fmt.Println("sliceA1:", sliceA1)
	fmt.Println("sliceA2:", sliceA2)
	fmt.Println("sliceA3:", sliceA3)

	// 通常のスライス式
	sliceB1 := []int{1, 2, 3, 4, 5}
	sliceB2 := sliceB1[0:2] // sliceB2の容量は5

	// sliceB2の容量に空きがあるため、
	// 要素追加時は既存のスライスを使用する
	// 既存のスライスを使用したため、sliceB1に影響する
	sliceB3 := append(sliceB2, 6)
	fmt.Println("sliceB1:", sliceB1)
	fmt.Println("sliceB2:", sliceB2)
	fmt.Println("sliceB3:", sliceB3)
}
