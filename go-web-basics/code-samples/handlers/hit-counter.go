package main

import (
	"fmt"
	"log"
	"net/http"
)

type Counter int

//This method makes a Counter satisfy the Handler interface
func (h *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*h++
	fmt.Fprintf(w, "This app's hit count: %d", *h)
}

//A handler function for the /sloths route
func slothsRule(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sloths rule!")
}

func main() {
	//Initialize our counter
	counter := Counter(0)

	mux := http.NewServeMux()
	//Use the counter as a Handler
	mux.Handle("/", &counter)

	//Use http.HandlerFunc to convert slothsRule to a Handler
	slothsRuleHandler := http.HandlerFunc(slothsRule)
	mux.Handle("/sloths", slothsRuleHandler)

	//Logs what URL the request is to and then sends the request to mux
	//by calling its ServeHTTP method.
	logAndServe := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to " + r.URL.String())
		mux.ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: logAndServe,
	}
	server.ListenAndServe()
}
