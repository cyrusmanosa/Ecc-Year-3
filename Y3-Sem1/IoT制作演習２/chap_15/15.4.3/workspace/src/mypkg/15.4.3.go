// パッケージの説明
package mypkg

import "fmt"

// FuncDoc関数の説明
func FuncDoc(n int) {
	// BUG(tobi):バグ用のコメント
	fmt.Println("funcDoc", n)
	// このコメントはドキュメント化されない
}
