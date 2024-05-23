package main

import "fmt"

// 型による条件分岐を確認する関数
func typeSwitch(i interface{}) {
	// iに格納している値の型で分岐する
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v) // vはint型
	case string:
		fmt.Println("string:", v) // vはstring型
	case int32, int64:
		fmt.Println("int32/64:", v) // vはiと同じ型(interface{}型)
	default:
		fmt.Println("default:", v) // vはiと同じ型(interface{}型)
	}
}

// 値がnilのときの条件分岐を確認する関数
func typeSwitchNil(i interface{}) {
	// iに格納している値の型で分岐する
	switch v := i.(type) {
	case nil: // 型のないnilのみ該当する
		fmt.Println("nil:", v)
	case *int: // *int型のnilはこちらに該当する
		fmt.Println("*int:", v)
	}
}

func main() {
	typeSwitch(1)        // 「int: 1」を出力
	typeSwitch("a")      // 「string: a」を出力
	typeSwitch(int32(1)) // 「int32/64: 1」を出力
	typeSwitch(1.1)      // 「default: 1.1」を出力

	typeSwitchNil(nil) // 「nil: nil」を出力
	var n *int = nil
	typeSwitchNil(n) // 「*int: nil」を出力
}
