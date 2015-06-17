package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Matches ONLY /sloths, not /sloths/ or /sloths/are-awesome
	mux.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sloths rule!")
	})

	//Matches /kangaroos, /kangaroos/, and /kangaroos/tree-kangaroos
	mux.HandleFunc("/kangaroos/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kangaroos for the win!")
	})

	//Since "/" ends with a slash, it matches all URL paths, so "/"
	//is the catch-all route.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lemurs get the catch-all route! Lemurs are where it's at!")
	})

	// tea/hibiscus matches this route, but the route below is longer
	mux.HandleFunc("/tea/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "One tea coming right up!")
	})

	//This route has a longer path so a request to tea/hibiscus matches this
	//route, not the one above.
	mux.HandleFunc("/tea/hibiscus", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "One hibiscus tea coming right up!")
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: mux,
	}
	server.ListenAndServe()
}
