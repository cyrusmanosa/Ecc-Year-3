// src1.go
// このソースファイル(src1.go)はpkg1パッケージに所属
package pkg1

// pkg1パッケージのf1 関数
func f1() {
	f2()  // 同じパッケージ内にあるf2関数を使用
}
