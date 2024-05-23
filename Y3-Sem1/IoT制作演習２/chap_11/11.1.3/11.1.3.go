package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 3) // 長さ2、容量3のスライス
	slice2 := append(slice1, 1) // 追加時に容量を超えない
	slice1[0] = 2               // slice2に影響する

	slice3 := make([]int, 2, 2) // 長さ2、容量2のスライス
	slice4 := append(slice3, 1) // 追加時に容量を超える
	slice3[0] = 2               // slice4に影響しない

	fmt.Println("len:", len(slice1), len(slice2), len(slice3), len(slice4))
	fmt.Println("cap:", cap(slice1), cap(slice2), cap(slice3), cap(slice4))
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)

	// slice1にslice2を追加
	slice5 := append(slice1, slice2...)
	fmt.Println("slice5:", slice5)

	// slice6にslice5をコピー
	slice6 := make([]int, 6)
	copy(slice6, slice5)
	fmt.Println("slice6:", slice6)
}
