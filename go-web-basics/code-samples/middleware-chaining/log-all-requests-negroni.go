package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	//simpleLog converted to a middleware function Negroni can use
	simpleLog := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.String())
	}

	serveMux := http.NewServeMux()

	//Routes
	serveMux.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir("public/images"))))
	serveMux.HandleFunc("/ducks", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/ducks.html")
	})
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Beware of ducks! Duck venom can turn people into ducks!")
	})

	stack := negroni.New() //Create a Negroni middleware stack with negroni.New

	//You can add an HTTP handler function to the middleware stack with
	//UseHandlerFunc; first we're adding the logger middleware to the stack.
	stack.UseHandlerFunc(simpleLog)

	//Add a handler like the router to the middleware stack with UseHandler.
	//In this case we're adding our ServeMux to the stack.
	stack.UseHandler(serveMux)

	//Use the Negroni middleware stack as our Server's Handler
	server := &http.Server{
		Addr:    ":1123",
		Handler: stack,
	}

	server.ListenAndServe()
}
