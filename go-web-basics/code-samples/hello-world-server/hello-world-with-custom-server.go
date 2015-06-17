package main

import (
	"fmt"
	"net/http"
)

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", serveHello)
	server := &http.Server{
		Addr:    ":1123",
		Handler: http.DefaultServeMux,
	}
	server.ListenAndServe()
}
