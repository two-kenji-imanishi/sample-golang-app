package main

import (
	"handlers/handlers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
