// pkg1/src1.go
package pkg1

import "../pkg2"    // パッケージpkg2をインポート(相対パス)

func F1() {
	pkg2.F2()       // パッケージpkg2の関数F2を使用
}
