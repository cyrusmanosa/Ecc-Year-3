package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"time"
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

	// Mmapを使用し、GPIO関連レジスタを[]uint8の変数に対応 … ③
	mem8, err = syscall.Mmap(
		int(file.Fd()), // ファイルディスクリプタ
		gpioBaseAddr,   // 割り当てるレジスタの開始アドレス
		blockSize,      // ブロックサイズの指定
		syscall.PROT_READ|syscall.PROT_WRITE, // 読み書き可能なモード
		syscall.MAP_SHARED)                   // マッピングを共有する

	// エラーの場合は終了
	if err != nil {
		return
	}

	// []uint8から[]uint32に変換 … ④
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&mem8)) // mem8のSliceHeaderを取得
	// LenおよびCapを4分の1にする
	header.Len /= 4 // (32 bit / 8bit = 4 なので4で割る)
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

	gpioRegister[0] = (gpioRegister[0] &^ (7 << 12)) // 入力にするためbit14からbit12に「000」を設定

	gpioRegister[37] = 2         // 「10」 = 2を設定することでプルアップ
	time.Sleep(time.Microsecond) // 1マイクロ秒待機
	gpioRegister[38] = 1 << 4    // GPIO4に対応するビット4に「1」を設定
	time.Sleep(time.Microsecond) // 1マイクロ秒待機

	isPushSwitch := false // 現在のスイッチ判定状態を保持する変数。初期値は「false」

	for {
		count := 0
		for i := 0; i < 3; i++ { // 3回ループし、判定する
			if gpioRegister[13]&(1<<4) == 0 {
				count += 1 // 「ON」と判定するとカウントアップ
			}
			time.Sleep(time.Millisecond) // 1ミリ秒待機
		}
		if isPushSwitch == false && count == 3 { // スイッチがOFF -> ONに変化するときに「ON」を表示
			isPushSwitch = true
			fmt.Println("ON")
		} else if isPushSwitch == true && count == 0 { // スイッチがON -> OFFに変化するときに「OFF」を表示
			isPushSwitch = false
			fmt.Println("OFF")
		}
	}

}
