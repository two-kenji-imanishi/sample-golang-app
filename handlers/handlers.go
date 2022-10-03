package handlers

import (
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to World!\n")
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello GAE World 2022!\n")
}
