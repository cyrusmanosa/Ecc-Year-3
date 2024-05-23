package main

import (
	"fmt"
	"net/http"
)

// 「http://localhost:4000/hello」の処理 … ①
func helloHandle(w http.ResponseWriter, r *http.Request) {
	h := `
<html><head><title>Hello</title></head><body>
  Hello
</body></html>
`
	fmt.Fprint(w, h)
}

// 「http://localhost:4000/goodbye」の処理 … ②
func goodbyeHandle(w http.ResponseWriter, r *http.Request) {

	h := `
<html><head><title>Goodbye</title></head><body>
  Goodbye
</body></html>
`
	fmt.Fprint(w, h)
}

func main() {
	// URLごとに関数を登録 … ③
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/goodbye", goodbyeHandle)

	// Webサーバを起動 … ④
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
