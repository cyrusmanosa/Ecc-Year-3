package main

import (
	"fmt"
	"html/template" // html/templateパッケージのインポートを追加 … ①
	"net/http"
)

// serverHTML型の宣言 … ②
type serverHTML struct {
	Title string
	Body  string
}

// テンプレートの設定 … ③
var tmpl = template.Must(template.ParseFiles("html.tmpl"))

// ServeHTTPメソッドの宣言 … ④
func (s *serverHTML) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HTTPヘッダの設定
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// テンプレートとデータを組み合わたHTMLをクライアントに送信する
	if err := tmpl.Execute(w, s); err != nil {
		panic(err)
	}
}

func main() {
	// URLごとに値を登録
	helloHTML := &serverHTML{"Hello Title", "Hello"}
	goodbyeHTML := &serverHTML{"Goodbye Title", "Goodbye"}
	http.Handle("/hello", helloHTML)
	http.Handle("/goodbye", goodbyeHTML)

	// Webサーバを起動
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
