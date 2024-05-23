package myapp

import (
	"appengine"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", mainHandle)
}
func mainHandle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id := appengine.RequestID(c)
	fmt.Fprint(w, "Request ID: "+id)
}
