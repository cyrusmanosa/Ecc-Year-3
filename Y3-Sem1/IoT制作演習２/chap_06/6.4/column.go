package main

import (
	"fmt"
	"math"
)

func main() {
	var f1 float32 = 0.3
	var b1 uint32 = math.Float32bits(f1)
	fmt.Printf("%.24f\n", f1) // 0.300000011920928955078125
	fmt.Printf("%032b\n", b1) // 001111101001100110011001100110100

	f1 += 0.4
	var f2 float32 = 0.7
	b1 = math.Float32bits(f1)
	var b2 uint32 = math.Float32bits(f2)

	fmt.Printf("%.24f\n", f1) // 0.700000047683715820312500
	fmt.Printf("%.24f\n", f2) // 0.699999988079071044921875
	fmt.Printf("%032b\n", b1) // 00111111001100110011001100110100
	fmt.Printf("%032b\n", b2) // 00111111001100110011001100110011
	fmt.Println(f1 == f2)     // false

	var f3 float32 = 0.3 + 0.4
	var b3 uint32 = math.Float32bits(f3)
	fmt.Printf("%.24f\n", f3) // 0.699999988079071044921875
	fmt.Printf("%032b\n", b3) // 00111111001100110011001100110011
	fmt.Println(f1 == f3)     // false
}
