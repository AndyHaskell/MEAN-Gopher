package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Serve a message raising awareness of the dangers of duck venom
	ducksHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Beware of ducks! Duck venom can turn people into ducks!")
	})

	//Logs the request's HTTP verb and URL and then calls ducksHandler's
	//ServeHTTP
	simpleLog := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.String()) //Log the request
		ducksHandler.ServeHTTP(w, r) //Now pass the request to ducksHandler
	})

	//Middleware chaining function. Takes in a Handler and returns a middleware
	//that when given a request logs the request's HTTP verb and URL and then
	//runs the Handler passed in
	logRequest := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s - %s\n", r.Method, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}

	logRequestAndRaiseDuckVenomAwarenessChain := logRequest(ducksHandler)

	http.Handle("/", ducksHandler)
	http.Handle("/simplelog", simpleLog)
	http.Handle("/duckChain", logRequestAndRaiseDuckVenomAwarenessChain)
	http.ListenAndServe(":1123", http.DefaultServeMux)
}
