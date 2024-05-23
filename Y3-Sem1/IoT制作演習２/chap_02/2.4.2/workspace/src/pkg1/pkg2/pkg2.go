// $GOPATH/src/pkg1/pkg2
package pkg2

// internalの親ディレクトリ内のパッケージからはインポート可能
import _ "pkg1/internal/pkg3" // ②
