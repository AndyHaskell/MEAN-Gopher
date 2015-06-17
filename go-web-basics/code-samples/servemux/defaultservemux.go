package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.DefaultServeMux.HandleFunc("/sloths",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Sloths rule!")
		})

	// This is equivalent to:
	//
	// http.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "Sloths rule!")
	// })

	http.ListenAndServe(":1123", nil)
}
