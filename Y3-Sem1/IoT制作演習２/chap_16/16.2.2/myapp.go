package myapp

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// エンティティとして読み書きする構造体
type Item struct {
	Name  string
	Count int
}

// テンプレートに使用する構造体
type ViewData struct {
	KeyName string
	Op      string
	Item
}

// 初期化処理
func init() {
	http.HandleFunc("/", mainHandle)
}

// テンプレートファイルの読み込み
var tmpl = template.Must(template.ParseFiles("html.tmpl"))

func mainHandle(w http.ResponseWriter, r *http.Request) {
	// 入力フォームの読み込み
	keyName := r.FormValue("key_name")
	op := r.FormValue("op")
	name := r.FormValue("name")
	count, _ := strconv.Atoi(r.FormValue("count"))

	// キーの作成
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Item", keyName, 0, nil)

	// テンプレート用データの作成
	viewData := ViewData{KeyName: keyName}

	// エンティティの操作
	switch op {
	case "put":
		// エンティティの追加・更新
		item := Item{name, count}
		if _, err := datastore.Put(c, key, &item); err != nil {
			log.Println(err)
		}

		viewData.Item = item
		viewData.Op = "追加・更新"

	case "get":
		// エンティティの読み込み
		item := Item{}
		if err := datastore.Get(c, key, &item); err != nil {
			log.Println(err)
		}

		viewData.Item = item
		viewData.Op = "検索"

	case "delete":
		// エンティティの削除
		if err := datastore.Delete(c, key); err != nil {
			log.Println(err)
		}

		viewData.Op = "削除"
	}

	// クライアントへの送信
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, &viewData); err != nil {
		log.Println(err)
	}
}
