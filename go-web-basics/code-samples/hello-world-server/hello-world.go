package main

import (
    "fmt"
    "net/http"
)

func serveHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    http.HandleFunc("/", serveHello)
    http.ListenAndServe(":1123", nil)
}