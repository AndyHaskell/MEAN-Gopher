package main

import "net/http"

//HelloWorld as a global Handler
var HelloWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
})

//serveHelloWorld as a handler function
func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

//initialize the ServeMux
func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", serveHelloWorld)
	return mux
}

func main() {
	mux := InitRouter()
	svr := &http.Server{
		Addr:    ":1123",
		Handler: mux,
	}
	svr.ListenAndServe()
}
