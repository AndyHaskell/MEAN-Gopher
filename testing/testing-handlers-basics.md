# Basics of testing `net/http` Handlers

## Unit testing your handlers

**Basic idea**: When a `Handler` gets a request, the `Handler`'s `ServeHTTP` method is called, so to simulate a request being handled, **make a `*Request` and `ResponseWriter` and call your `Handler`'s `ServeHTTP` on those.**

### Getting a Request

A `Request` is a struct and to make one, use `NewRequest`

```go
http.NewRequest(method, url string, body io.Reader)
```

where the method is what HTTP verb to use for the request (**GET**, **POST**, etc), the URL is the request's path (ex `"/sloths"`, `"/kangaroos/tree-kangaroos"`, or even just `""`), and the body is an optional `io.Reader`.

For example:

```
req := http.NewRequest("GET", "", nil)
```

makes a GET request with no request URL path and no body while

```
req := http.NewRequest("GET", "/is-fibonacci/21")
```

makes a GET request to the path `/is-fibonacci/21` with no body.

### Getting a ResponseWriter

Unlike a `Request`, which is a struct, a `ResponseWriter` is an interface, so to get a `ResponseWriter` we need to make an object that implements the interface.

Conveniently, Go has a standard **`net/http/httptest`** library made for testing web apps, which has a `ResponseRecorder` type that implements `ResponseWriter` and is made for testing and keeping track of responses. To get one, call

```go
w := httptest.NewRecorder()
```


## Unit testing a Handler

Let's say we have a handler that just serves a Hello world message:

```go
var HelloWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
})
```

To make a test function testing this Handler we would first initialize our `ResponseRecorder` and `Request` with:

```go
w := httptest.NewRecorder()
r, err := http.NewRequest("GET", "", nil)
```

Test the request by calling:

```go
HelloWorld.ServeHTTP(w, r)
```

And then test the status code and the contents of the response with these if statements:

```go
if w.Code != 200 {
    t.Fatalf("Response status code expected 200, got %d", w.Code)
}
if w.Body.String() != "Hello, world!" {
    t.Fatalf("w.Body.String() failed, expected %v, got %v",
        "Hello, world!", w.Body.String())
}
```

## Testing a Handler in a server

To see if your `Handler`s work as part of a full-fledged server and plugged into stuff like a middleware stack, `httptest` provides its own `Server` type.

Passing any `Handler` into `httptest.NewServer` creates a test server.  Like with a `Server` from `net/http`, that `Handler` can be anything from a basic handler function to a middleware chain to an HTTP router like a `ServeMux` or Gorilla `Router`. For example:

```go
svr := httptest.NewServer(HelloWorld)
```

Or say we had `HelloWorld` as part of a `ServeMux`:

```go
var Mux = http.NewServeMux()
Mux.Handle("/hello", HelloWorld)
```

Then in a test file like `server_test.go`, we would start the server with:

```go
svr := httptest.NewServer(Mux)
```

Which both creates and starts a test server that uses our `ServeMux` as its `Handler`. You can then send HTTP requests to the server with either a `Client` or from `net/http`'s built-in methods like `Get` and `Post`. For example,

```go
res, err := http.Get(svr.URL+"/hello")
if err != nil {
    t.Fatalf(err.Error())
}

responseText, err := ioutil.ReadAll(res.Body)
if err != nil {
    t.Fatalf(err.Error())
}
if string(responseText) != "Hello, world!" {
    t.Fatalf("string(responseText) failed, expected %v, got %v",
        "Hello, world!", string(responseText))
}
```

`svr.URL` is the server's URL; the server generates its own port for you. So the line

```go
res, err := http.Get(svr.URL+"/hello")
```

sends a GET request to the server, which gives us back a `net/http` `Response`. The contents of its body are then analyzed with `ioutil.ReadAll(res.Body)`.

Note that in order to do tests on both `Handler`s and on larger structures that use them, it's a good idea to write modularized Go packages to maximize testability. For example, having a function that generates your router...

```go
//serveHelloWorld as a handler function
func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

//initialize the ServeMux
func InitServeMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", serveHelloWorld)
    return mux
}
```

Allows you to use that router in both your real server and httptest server without depending on extra global variables.
