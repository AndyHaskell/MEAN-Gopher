package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	serveOrderForm := func(w http.ResponseWriter, r *http.Request) {
		//Restrict the route to only GET requests
		if r.Method == "GET" || r.Method == "" {
			http.ServeFile(w, r, "pages/order-form.html")
		} else {
			http.Error(w, "405 Method Not Allowed", 405)
		}
	}

	serveSendOrder := func(w http.ResponseWriter, r *http.Request) {
		//Restrict the route to only POST requests
		if r.Method == "POST" {
			//Parse the POST data with r.ParseForm() and
			//get the data with r.Form.Get()
			r.ParseForm()
			beverage := html.EscapeString(r.Form.Get("beverage"))
			name := html.EscapeString(r.Form.Get("name"))

			fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
				beverage, name)
		} else {
			http.Error(w, "405 Method Not Allowed", 405)
		}
	}
	mux.HandleFunc("/order-form", serveOrderForm)
	mux.HandleFunc("/send-order", serveSendOrder)

	//This is the version of /coffee-shop that isn't modularized
	/*mux.HandleFunc("/coffee-shop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "" {
			http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
		} else if r.Method == "POST" {
			r.ParseForm()
			beverage := html.EscapeString(r.Form.Get("beverage"))
			name := html.EscapeString(r.Form.Get("name"))

			fmt.Fprintf(w, "One %s coming right up, %s!", beverage, name)
		} else {
			http.Error(w, "405 Method Not Allowed", 405)
		}
	})*/

	orderForm := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
	})
	sendOrder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		beverage := html.EscapeString(r.Form.Get("beverage"))
		name := html.EscapeString(r.Form.Get("name"))

		fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
			beverage, name)
	})

	//This is the modularized version of coffee-shop
	mux.HandleFunc("/coffee-shop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "" {
			orderForm.ServeHTTP(w, r)
		} else if r.Method == "POST" {
			sendOrder.ServeHTTP(w, r)
		} else {
			http.Error(w, "405 Method Not Allowed", 405)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqMethod := r.Method
		if reqMethod == "" {
			reqMethod = "GET"
		}

		fmt.Fprintf(w, "Your request method is %s", reqMethod)
	})

	server := &http.Server{
		Addr:    ":1123",
		Handler: mux,
	}
	server.ListenAndServe()
}
