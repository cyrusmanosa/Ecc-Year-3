package main

import "fmt"

func main() {
	var (
		f32r float32 = 1.1
		f32i float32 = 2.2
		f64r float64 = 1.1
		f64i float64 = 2.2
	)

	// complex関数
	// float32型の引数から、complex64型の複素数を取得
	var c64 complex64 = complex(f32r, f32i)
	fmt.Println("c64: ", c64)

	// float64型の引数から、complex128型の複素数を取得
	var c128 complex128 = complex(f64r, f64i)
	fmt.Println("c128:", c128)

	// real関数、imag関数
	// complex64型の引数から、float32型の実数と虚数を取得
	f32r = real(c64)
	f32i = imag(c64)
	fmt.Println("f32r:", f32r)
	fmt.Println("f32i:", f32i)

	// complex128型の引数から、float64型の実数と虚数を取得
	f64r = real(c128)
	f64i = imag(c128)
	fmt.Println("f64r:",f64r)
	fmt.Println("f64i:",f64i)
}
