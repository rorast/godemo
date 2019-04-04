package main

import (
	"io"
	"net/http"
)

func fistPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is my fist page!中文字</h1>")
}

func main() {
	http.HandleFunc("/", fistPage)
	http.ListenAndServe(":8000", nil)

}
