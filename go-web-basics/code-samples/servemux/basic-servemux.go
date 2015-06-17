package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sloths rule!")
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: mux,
	}
	server.ListenAndServe()
}
