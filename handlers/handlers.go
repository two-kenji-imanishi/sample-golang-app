package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to World!\n")
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello! My name is %s.", os.Getenv("NAME")))
}
