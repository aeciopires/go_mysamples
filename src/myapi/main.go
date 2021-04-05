package main

import (
	"log"
	"net/http"
)

func main() {
	listen := "localhost:8080"
	router := NewRouter()

	log.Fatal(http.ListenAndServe(listen, router))
}
