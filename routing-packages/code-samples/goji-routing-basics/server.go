package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

//A simple chain of Goji handler functions
func youreNo1000000(c web.C, w http.ResponseWriter, r *http.Request) {
	c.Env["hitNumber"] = 1000000
	serveHitNumber(c, w, r)
}
func serveHitNumber(c web.C, w http.ResponseWriter, r *http.Request) {
	hitNumber := c.Env["hitNumber"]
	fmt.Fprintf(w, "You're totally viewer number %d!", hitNumber)
}

func main() {
	//Initialize the router with the EnvInit middleware
	m := web.New()
	m.Use(middleware.EnvInit)

	//Make a plain path
	m.Handle("/sloths", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sloths rule!")
	})

	//Route parameters
	m.Handle("/:flavor/tea", func(c web.C, w http.ResponseWriter, r *http.Request) {
		flavor := c.URLParams["flavor"]
		fmt.Fprintf(w, "I could go for some %s tea!", flavor)
	})

	//Path prefix with a *
	m.Handle("/img/*", http.StripPrefix("/img/",
		http.FileServer(http.Dir("public/images"))))

	//GET-specific route
	m.Get("/get-route", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This route only responds to GET requests")
	})

	//Route with a regular expression path
	coffeeRegexp := regexp.MustCompile(`^/(coffee)+$`)
	m.Get(coffeeRegexp, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lemurs = sloths that had too much coffee")
	})

	m.Handle("/", youreNo1000000)

	//Catch-all route
	m.Handle("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This route matches all requests.")
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: m,
	}
	server.ListenAndServe()
}
