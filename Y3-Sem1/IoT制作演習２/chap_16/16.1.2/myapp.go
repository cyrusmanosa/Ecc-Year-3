package myapp // パッケージ名 … ①

import (
	"fmt"
	"net/http"
)

// init関数 … ②
func init() {
	// URLに関数を登録
	http.HandleFunc("/", mainHandle)
	// ListenAndServe関数の呼び出しは不要
}

// mainHandle関数
func mainHandle(w http.ResponseWriter, r *http.Request) {
	// クライアント（ブラウザ）に送信
	fmt.Fprint(w, "Hello")
}
