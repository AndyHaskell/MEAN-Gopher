package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test the hello world handler by itself
func TestHelloWorld(t *testing.T) {
	//Make the ResponseRecorder and Request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	//Have HelloWorld handle the request with its ServeHTTP
	HelloWorld.ServeHTTP(w, r)

	//Test the status code and the response body
	if w.Code != 200 {
		t.Fatalf("Response status code expected 200, got %d", w.Code)
	}
	if w.Body.String() != "Hello, world!" {
		t.Fatalf("w.Body.String() failed, expected %v, got %v",
			"Hello, world!", w.Body.String())
	}
}

//Test the hello world handler as part of a server
func TestHelloWorldServer(t *testing.T) {
	//Initialize the router and server
	mux := InitRouter()
	svr := httptest.NewServer(mux)

	//Request the hello world route
	res, err := http.Get(svr.URL + "/hello")
	if err != nil {
		t.Fatalf(err.Error())
	}

	//Get the response's body text and verify that it reads "Hello, world!"
	responseText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if string(responseText) != "Hello, world!" {
		t.Fatalf("string(responseText) failed, expected %v, got %v",
			"Hello, world!", string(responseText))
	}
}
