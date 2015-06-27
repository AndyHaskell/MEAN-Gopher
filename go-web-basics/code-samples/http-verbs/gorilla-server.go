package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
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

	m := mux.NewRouter() //Create a Gorilla mux Router

	//You restrict routes to specific HTTP methods with Route.Methods()
	m.HandleFunc("/order-form", serveOrderForm).Methods("GET")
	m.HandleFunc("/send-order", serveSendOrder).Methods("POST")

	//You can also register multiple handlers to the same path in Gorilla and
	//having Gorilla resolve which one to serve with Route.Methods
	m.Path("/coffee-shop").HandlerFunc(serveCoffeeShopOrderForm).Methods("GET")
	m.Path("/coffee-shop").HandlerFunc(serveSendOrder).Methods("POST")

	m.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
