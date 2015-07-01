package main

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
)

func main() {
	logRequest := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s - %s\n", r.Method, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()

	//Routes
	mux.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir("public/images"))))
	mux.HandleFunc("/ducks", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/ducks.html")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Beware of ducks! Duck venom can turn people into ducks!")
	})

	//An Alice middleware chain that chains logRequest with the ServeMux
	logAndServeChain := alice.New(logRequest).Then(mux)

	server := &http.Server{
		Addr:    ":1123",
		Handler: logAndServeChain,
	}

	server.ListenAndServe()
}
