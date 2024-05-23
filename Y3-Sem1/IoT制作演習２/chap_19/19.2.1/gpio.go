package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

const (
	baseAddr     = 0x3F000000          // Raspberry Piのレジスタ開始アドレス
	gpioBaseAddr = baseAddr + 0x200000 // GPIO関連レジスタの開始アドレス
	blockSize    = 4096                // ブロックサイズ
)

var (
	mem8         []uint8
	gpioRegister []uint32
)

func allocateRegister() (err error) {
	// /dev/memをfile変数に割り当てる … ①
	file, err := os.OpenFile(
		"/dev/mem",          // 開くファイルの指定
		os.O_RDWR|os.O_SYNC, // アクセス権。読み書き＋同期モード
		0) // ファイルモードは特に指定しない

	// エラーの場合は終了
	if err != nil {
		return
	}

	// この関数を抜けるときにファイルをクローズする
	defer file.Close()

	// Mmapを使用し、GPIO関連レジスタを[]uint8の変数にマッピング … ②
	mem8, err = syscall.Mmap(
		int(file.Fd()), // ファイルディスクリプタ
		gpioBaseAddr,   // 割り当てるレジスタの先頭アドレス
		blockSize,      // ブロックサイズの指定
		syscall.PROT_READ|syscall.PROT_WRITE, // 読み書き可能なモード
		syscall.MAP_SHARED)                   // マッピングを共有する

	// エラーの場合は終了
	if err != nil {
		return
	}

	// []uint8から[]uint32に変換 … ③
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&mem8)) // mem8のSliceHeaderを取得
	// LenおよびCapを4分の1にする
	header.Len /= 4 // (32ビット / 8ビット = 4 なので4で割る)
	header.Cap /= 4

	gpioRegister = *(*[]uint32)(unsafe.Pointer(&header)) // mem8のSliceHeaderを[]uint32にキャスト

	return nil

}

func main() {
	err := allocateRegister() // GPIOレジスタを変数に割り当て
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%032b\n", &gpioRegister[0])
}
