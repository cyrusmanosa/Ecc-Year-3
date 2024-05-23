// $GOPATH/src/pkg1
package pkg1

// internalの親ディレクトリ（"pkg1"）からはインポート可能
import _ "pkg1/internal/pkg3" // ①
