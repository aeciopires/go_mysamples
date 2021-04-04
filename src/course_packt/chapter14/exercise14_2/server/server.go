package main

import (
	"log"
	"net/http"
)

type server struct{}

// Func ServeHTTP serve page with JSON content
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "{\"message\": \"hello world\"}"
	w.Write([]byte(msg))
}

func main() {
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, server{}))
}
