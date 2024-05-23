package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// ポインタ型からunsafe.Pointer型への変換
	pointerInt1 := new(int)
	unsafePointer1 := unsafe.Pointer(pointerInt1)
	fmt.Println(unsafePointer1)

	// unsafe.Pointer型からuintptr型への変換
	u := uintptr(unsafePointer1)
	u += 2 // uintptr型は演算可能

	// uintptr型からunsafe.Pointer型への変換
	unsafePointer1 = unsafe.Pointer(u)
	fmt.Println(unsafePointer1)

	// unsafe.Pointer型から別のポインタ型への変換
	pointerInt2 := new(int)
	unsafePointer2 := unsafe.Pointer(pointerInt2)
	pointerString := (*string)(unsafePointer2)
	// 「pointerString := (*string)(pointerInt2pInt)  」はコンパイル時にエラーとなる

    s := *pointerString // この行でエラーとなる
    fmt.Println(s)
}
