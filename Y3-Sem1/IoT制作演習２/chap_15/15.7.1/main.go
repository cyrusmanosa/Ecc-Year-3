package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// Windowsの場合
	dll := "msvcrt.dll"

	// Mac OS X、Linuxの場合
	// dll := "libc.so"

	// 共有ライブラリにあるputs関数の読み込み
	mod := syscall.NewLazyDLL(dll)
	proc := mod.NewProc("puts")

	// 文字列を*byte型に変換
	s := syscall.StringBytePtr("call shared library")

	// *byte型をuintptr型に変換
	u := uintptr(unsafe.Pointer(s))

	// puts関数の読み出し
	proc.Call(u)
}
