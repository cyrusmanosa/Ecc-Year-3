package main

import (
	"fmt"
	"net/http"
)

// ServeHTTPメソッド用の構造体型
type Server struct{}

// httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HTMLの文字列 … ①
	h := `
<html><head><title>Hello</title></head><body>
  Hello
</body></html>
`

	// クライアント(ブラウザ)にHTMLを送信 … ②
	fmt.Fprint(w, h)
}

func main() {
	// Webサーバを起動 … ③
	http.ListenAndServe(":4000", Server{})
}
