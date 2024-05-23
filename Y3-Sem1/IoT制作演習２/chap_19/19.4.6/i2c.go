package main

import (
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"time"
)

func main() {
	// I2Cバス1を開く…①
	bus := embd.NewI2CBus(1)

	// ADT7410を連続変換モード、分解能13ビットに設定を変更…②
	bus.WriteByteToReg(0x48, 0x03, 0x00) // スレーブ0x48のレジスタアドレス0x03に0x00を書き込む
	time.Sleep(time.Second)              // モード変更後は240ms以上待つ

	// 温度値レジスタからデータを読み込む…③
	tempReg, _ := bus.ReadWordFromReg(0x48, 0x00) // スレーブ0x48のレジスタアドレス0x00から2バイト読み込む
	fmt.Printf("%04x\n", tempReg)

	// 温度値をセ氏温度に変換（分解能13ビットの場合）…④
	temp := float32(uint16(tempReg) >> 3) / 16.0 // 上位13ビットを使い16で割ることでセ氏温度に変換
	fmt.Printf("%f\n", temp)

	// ADT7410を1 SPSモード、分解能16ビットに設定を変更…⑤
	bus.WriteByteToReg(0x48, 0x03, 0xC0) // スレーブ0x48のアドレス0x03に0xC0を書き込む
	time.Sleep(time.Second)              // モード変更後は60ms以上待つ

	// 温度値レジスタからデータを読み込む…⑥
	tempReg, _ = bus.ReadWordFromReg(0x48, 0x00) //0x00から2バイト取得
	fmt.Printf("%04x\n", tempReg)

	// 温度値をセ氏温度に変換（分解能16ビットの場合）…⑦
	temp = float32(tempReg) / 128.0 // 128で割ることでセ氏温度に変換
	fmt.Printf("%f\n", temp)

}
