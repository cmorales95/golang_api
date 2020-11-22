package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("server is run in localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}

