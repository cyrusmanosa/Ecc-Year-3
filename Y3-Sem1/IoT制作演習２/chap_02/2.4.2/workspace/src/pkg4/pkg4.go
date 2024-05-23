// $GOPATH/src/pkg4
package pkg4

// internalの外側のパッケージからはインポート不可能
// 下記のインポート宣言はコンパイル時にエラーとなる
import _ "pkg1/internal/pkg3" // ③
