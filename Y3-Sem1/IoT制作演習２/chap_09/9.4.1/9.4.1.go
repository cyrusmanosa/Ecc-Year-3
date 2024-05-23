package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONデータに変換する構造体 … ①
type JsonData struct {
	A string
	B string `json:"code"`
	C int    `json:",string"`
	D string `json:",omitempty"`
	E string `json:"-"`
}

// 「http://localhost:4000/json」の処理 … ②
func jsonHandler(w http.ResponseWriter, r *http.Request) {

	// 構造体をJSONデータに変換して出力
	data := JsonData{A: "データ1", B: "データ2", C: 2, D: "", E: "データ5"}
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// jは[]byte型であるため、文字列型に変換。
	fmt.Fprint(w, string(j))
}

// 「http://localhost:4000/html」の処理 … ③
func htmlHandler(w http.ResponseWriter, r *http.Request) {

	// 非同期にJSONデータを取得するHTMLとJavaScript
	h := `
<html><head><title>JSON</title></head><body>
<a href="#" onclick="reqJson()">JSON取得</a>
<script>
function reqJson() {
	var xhr = new XMLHttpRequest();
	xhr.open("GET", "/json");
	xhr.onload = function() { alert(xhr.responseText) }
	xhr.send();
}
</script>
</body></html>
`

	fmt.Fprint(w, h)
}

func main() {
	// URLに関数を登録
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/html", htmlHandler)

	// Webサーバ起動
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
