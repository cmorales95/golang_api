package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// register	route and handler
	http.HandleFunc("/sayhello", sayhello)
	log.Println("app is running in localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello World")
}