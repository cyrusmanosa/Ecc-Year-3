package main

import (
	"fmt"
	"net/http"
)

// serverHTML構造体型  の宣言 … ①
type serverHTML struct {
	title string
	body  string
}

// ServeHTTPメソッドの宣言 … ②
func (s *serverHTML) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := `
<html><head><title>` + s.title + `</title></head><body>
` + s.body + `
</body></html>
`
	fmt.Fprint(w, h)
}

func main() {
	// URLごとに値を登録 … ③
	helloHTML := &serverHTML{"Hello Title", "Hello"}
	goodbyeHTML := &serverHTML{"Goodbye Title", "Goodbye"}
	http.Handle("/hello", helloHTML)
	http.Handle("/goodbye", goodbyeHTML)

	// Webサーバを起動
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
