package main

import (
	"fmt"
	"html"
	"math/big" // math/bigパッケージのインポートを追加 … ①
	"net/http"
)

// ServeHTTPメソッド用の構造体型
type Server struct{}

// httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// フォームの入力値を取得 … ②
	left := r.FormValue("left")   // 左項目（テキストボックスの値）
	right := r.FormValue("right") // 右項目（テキストボックスの値）
	op := r.FormValue("op")       // 演算子（ラジオボタンの値）

	// 文字列を整数に変換 … ②
	leftInt := &big.Int{}
	rightInt := &big.Int{}
	_, leftOK := leftInt.SetString(left, 10)
	_, rightOK := rightInt.SetString(right, 10)

	// 四則演算の処理 … ③
	// 変換エラーがなければ、演算子に従って計算
	var result string
	if leftOK && rightOK {
		resultInt := &big.Int{}
		// 演算子ごとに分岐
		switch op {
		case "add":     // 加算
			resultInt.Add(leftInt, rightInt)
		case "sub":     // 減算
			resultInt.Sub(leftInt, rightInt)
		case "multi":   // 乗算
			resultInt.Mul(leftInt, rightInt)
		case "div":     // 除算
			resultInt.Div(leftInt, rightInt)
		}
		result = resultInt.String()
	}

	// HTMLの文字列 … ⑤
	h := `
<html><head><title>電卓アプリ</title></head><body>
  <form action="/" method="post">
	左項目：<input type="text" name="left"><br>
	右項目：<input type="text" name="right"><br>
	演算子：
	<input type="radio" name="op" value="add" checked> ＋
	<input type="radio" name="op" value="sub"> −
	<input type="radio" name="op" value="multi"> ×
	<input type="radio" name="op" value="div"> ÷
	<br><input type="submit" name="送信"><hr>

	[フォームの入力値]<br>
	左項目：` + html.EscapeString(left) + `<br>
	右項目：` + html.EscapeString(right) + `<br>
	演算子：` + html.EscapeString(op) + `<hr>
	演算結果：` + html.EscapeString(result) + `
  </form>
</body></html>
`

	// クライアント(ブラウザ)にHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	// Webサーバを起動
	http.ListenAndServe(":4000", Server{})
}
