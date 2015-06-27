package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()

	//Plain router have the same syntax as in net/http
	m.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sloths rule!")
	})

	//Use Router.PathPrefix for matching routes with prefixes
	m.PathPrefix("/img/").Handler(
		http.StripPrefix("/img/", http.FileServer(http.Dir("public/images"))))

	//Gorilla mux routes can take in route parameters with curly braces
	m.HandleFunc("/{flavor}/tea", func(w http.ResponseWriter, r *http.Request) {
		routeParams := mux.Vars(r)
		fmt.Fprintf(w, "I could go for some %s tea!", routeParams["flavor"])
	})

	//Regular expression routes in Gorilla mux are a slight variation on
	//route parameters.
	m.HandleFunc(`/{drink:(coffee)+}`,
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Lemurs = sloths that had too much coffee")
		})

	//Router.PathPrefix("/") creates a catch-all route
	m.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This route matches all requests.")
	})

	//A Gorilla mux Router is a Handler so we can use it as our Server's
	//main Handler.
	server := &http.Server{
		Addr:    ":1123",
		Handler: m,
	}
	server.ListenAndServe()
}
