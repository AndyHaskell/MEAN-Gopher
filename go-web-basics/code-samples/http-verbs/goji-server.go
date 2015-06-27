package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/zenazn/goji/web"
)

func main() {
	serveOrderForm := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/order-form.html")
	}
	serveCoffeeShopOrderForm := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
	}

	serveSendOrder := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		beverage := html.EscapeString(r.Form.Get("beverage"))
		name := html.EscapeString(r.Form.Get("name"))

		fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
			beverage, name)
	}

	m := web.New() //Create a Goji Mux

	//A Goji Mux comes with Get and Post methods for registering routes for
	//specific HTTP methods
	m.Get("/order-form", serveOrderForm)
	m.Post("/send-order", serveSendOrder)

	//You can register handlers to the same route with Get and Post
	m.Get("/coffee-shop", serveCoffeeShopOrderForm)
	m.Post("/coffee-shop", serveSendOrder)

	m.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		reqMethod := r.Method
		if reqMethod == "" {
			reqMethod = "GET"
		}

		fmt.Fprintf(w, "Your request method is %s", reqMethod)
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: m,
	}
	server.ListenAndServe()
}
