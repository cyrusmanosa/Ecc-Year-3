package main

import (
	"encoding/json"
	"fmt"
)

// JSONデータに変換する構造体
type JsonData struct {
	A string
	B int
}

func unmarshal(jsonText string) {
	// JSON形式の文字列を構造体に変換する
	var data JsonData
	err := json.Unmarshal(([]byte)(jsonText), &data)

	// JSON文法エラー
	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		fmt.Println("[文法エラー]", syntaxErr)
		fmt.Println("  Offset:", syntaxErr.Offset)
		return
	}

	// JSON型エラー
	if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		fmt.Println("[型エラー]", typeErr)
		fmt.Println("  Offset:", typeErr.Offset)
		fmt.Println("  Value:", typeErr.Value)
		fmt.Println("  Type:", typeErr.Type)
		return
	}

	// その他のエラー
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[変換OK]")
	fmt.Println(" ", data)
}

func main() {
	// JSON形式の文字列を構造体に変換する
	unmarshal(`{"A":"X", "B":1}`) // 正常
	unmarshal(`{"A":"X", "B":1]`) // 文法エラー
	unmarshal(`{"A":2,   "B":1}`) // 型エラー
}
