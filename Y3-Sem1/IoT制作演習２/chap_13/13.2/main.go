package main

import (
	"html/template"
	"net/http"
	"strconv" // strconvパッケージのインポートを追加 … ①
)

// 盤面などで使用する文字列
const maru, batsu = "○", "×"

// Board型の宣言
type Board [3][3]string

// ViewData型の宣言
type ViewData struct {
	Turn  string // 手番
	Board *Board // 盤面
}

// テンプレートの設定
var tmpl *template.Template = template.Must(template.ParseFiles("game.tmpl"))

// Executeメソッドの宣言
func (v *ViewData) Execute(w http.ResponseWriter) {
	// HTMLをクライアント（ブラウザ）に送信する
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, v); err != nil {
		panic(err)
	}
}

// gameHandle関数の宣言 … ②
func gameHandle(w http.ResponseWriter, r *http.Request) {
	_, nextTurn := turnFormValue(r) // 手番の入力値を取得する
	board := boardFormValue(r)      // 盤面の入力値を取得する

	// 値を設定してHTMLを送信する
	v := ViewData{nextTurn, board}
	v.Execute(w)
}

func main() {
	http.HandleFunc("/game", gameHandle)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}

// boardFormValue関数の宣言（盤面の値を取得） … ③
func boardFormValue(r *http.Request) *Board {
	var board Board
	for row, rows := range board {
		for col, _ := range rows {
			// 盤面のname属性「c00」～「c22」を作成
			name := "c" + strconv.Itoa(row) + strconv.Itoa(col)
			// 盤面の各項目を取得
			board[row][col] = r.FormValue(name)
		}
	}
	return &board
}

// 変数nextTurnMapに、次の手番を取得するマップを設定 … ④
var nextTurnMap = map[string]string{
	maru:  batsu,
	batsu: maru,
	"":    maru, // 「""」の場合、ゲーム開始時として「"○"」を取得
}

// turnFormValue関数の宣言（手番の値を取得） … ⑤
func turnFormValue(r *http.Request) (string, string) {
	turn := r.FormValue("turn")   // 現在の手番を取得
	nextTurn := nextTurnMap[turn] // マップを使用して次の手番を取得
	return turn, nextTurn
}
