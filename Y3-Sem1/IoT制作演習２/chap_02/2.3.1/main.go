// main.go
package main

import "./pkg1"     // パッケージpkg1をインポート(相対パス)

func main() {
	pkg1.F1()       // パッケージpkg1の関数F1を使用
}
