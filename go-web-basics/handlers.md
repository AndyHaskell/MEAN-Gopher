# All about Handlers
A `Handler` is the fundamental type for HTTP request handling in `net/http`. The definition of a `Handler` is any type with the method:

`ServeHTTP(w http.ResponseWriter, r *http.Request)`.

Any time a `Handler` gets an HTTP request, it calls its `ServeHTTP` method to handle it.

## A hit counter Handler

### In Node.js:

```javascript
var app = express()
var counter = 0

app.use(function(req, res) {
  counter++
  res.send("This app's hit count: " + counter)
})

http.createServer(app).listen(1123)
```

### In Go:

```go
type Counter int

//This method makes the Counter type satisfy the Handler interface
func (h *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    *h++
    fmt.Fprintf(w, "This app's hit count: %d", *h)
}

func main() {
    counter := Counter(0)

    mux := http.NewServeMux()
    mux.Handle("/", &counter)

    server := &http.Server {
        Addr: ":1123",
        Handler: mux,
    }
    server.ListenAndServe()
}
```

A `Handler` is registered to a `ServeMux`'s route with its `Handle` method, similar to how a handler function is registered with a its `HandleFunc` method.

Requests to this server are handled by `counter`, which increments itself and then sends a response saying how many hits the server has gotten.

In Node/Express, a handler is effectively a function with the signature `function(req, res)`, while in Go a handler is a type that has a `ServeHTTP` method. Both can have their own data besides the function for handling the requests, so both work effectively the same way.


## HandlerFunc

Since any type with a `ServeHTTP` method is a `Handler`, functions can be `Handler`s too. And if a function takes in a `ResponseWriter` and `*Request`, the function can be a `Handler` **where that function is used as its own `ServeHTTP` method!**

A function can be converted to a `Handler` doing just that with the `http.HandlerFunc` type. For example,

```go
slothsRuleHandler := http.HandlerFunc(slothsRule)
mux.Handle("/sloths", slothsRuleHandler)
```

Indeed, this is exactly how a `ServeMux`'s `HandleFunc` method works; it calls:

`yourServeMux.Handle(pattern, HandlerFunc(yourHandlerFunction))`,

making your handler function into a `Handler` where the function is its own `ServeHTTP` method.

## ServeMuxes are Handlers

A `Server` in Go always has a `Handler` property that takes in an `http.Handler`; all requests to a `Server` are processed by that `Handler`'s `ServeHTTP` method.

```go
mux := http.NewServeMux()
mux.Handle("/", &counter)

server := &http.Server {
    Addr: ":1123",
    Handler: mux,
}
server.ListenAndServe()
```

A `ServeMux`'s `ServeHTTP` sees matches a request to a route on the `ServeMux` and then **calls the `ServeHTTP` method for the `Handler` registered to that route**.

Since a `Server` can have **any** `Handler` for its `Handler` property, you could also pass in Handlers **that aren't `ServeMux`es**. For example, if you have a different router `Handler` like a router from the Gorilla mux package, you could do something like:

```go
router := mux.NewRouter()
router.Handle("/", &counter)

server := &http.Server {
    Addr: ":1123",
    Handler: router,
}
server.ListenAndServe()
```

and use the Gorilla router for all the routing instead of a `ServeMux`.

## Middleware chaining

A `Handler`'s `ServeHTTP` method can call another `Handler`'s `ServeHTTP` in the function body; if it can do that, that first `Handler` functions as middleware for the other `Handler`. For example, we could add a simple logger like this:

```go
//Logs what URL the request is to and then sends the request to mux
//by calling its ServeHTTP method.
logAndServe := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    log.Println("Request to " + r.URL.String())
    mux.ServeHTTP(w, r)
})

server := &http.Server {
    Addr: ":1123",
    Handler: logAndServe,
}
server.ListenAndServe()
```

What happens now is all requests go to `logAndServe` where the requested URL is logged and then `logAndServe` passes the request to our `ServeMux` **by calling its `ServeHTTP`**.

In Express that middleware would look like this:

```javascript
app.use(function(req, res, next){
  console.log('Request to', req.url)
  next()
})

/* ***Other Express routes*** */

http.createServer(app).listen(1123)
```

Notice that `ServeHTTP` in Go is used in the same way as `next` is used in Express; if a handler function can call another handler's `ServeHTTP`, it can be used as middleware the same way you would use middleware functions that call `next()` in Express.

Also notice that a `ServeMux` works by calling its `ServeHTTP`, which determines which route the request matches and then calls the `ServeHTTP` of the `Handler` on that route. That means a `ServeMux` isn't just a `Handler`, it's a Handler that functions as middleware!
