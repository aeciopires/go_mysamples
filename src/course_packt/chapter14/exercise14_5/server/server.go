package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}

	time.Sleep(10 * time.Second)
	msg := "hello client!"
	w.Write([]byte(msg))
}

func main() {
	listen := ":8080"
	log.Fatal(http.ListenAndServe(listen, server{}))
}
