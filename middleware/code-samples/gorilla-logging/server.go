package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the catch-all route")
	})

	//As our io.Writer, we use os.Stdout and we chain our LoggingHandler
	//with our ServeMux.
	logAndServe := handlers.LoggingHandler(os.Stdout, mux)

	//Now just pass in logAndServe as our server's Handler
	server := &http.Server{
		Addr:    ":1123",
		Handler: logAndServe,
	}
	server.ListenAndServe()
}
