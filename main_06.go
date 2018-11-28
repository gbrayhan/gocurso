package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Servidor Go")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080",nil)
}