# Middleware chaining in Go

## Recap on what Handlers are

In Go `net/http`, the definition of a `Handler` is any object with a `ServeHTTP` method with this signature:

`ServeHTTP(http.ResponseWriter, *http.Request)`

And a function can be converted to a `Handler` with the function `http.HandlerFunc`. That `Handler` calls that function in its `ServeHTTP`. For example:

```go
ducksHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Beware of ducks! Duck venom can turn people into ducks!")
})

http.Handle("/", ducksHandler)
```

For comparison, a handler function in Express is:

```javascript
var ducksHandler = function(req, res){
  res.send('Beware of ducks! Duck venom can turn people into ducks!')
}

app.use('/', ducksHandler)
```

So handler functions in Node.js work similarly to `Handler`s in Go and have the same parameters as a Go `ServeHTTP` method.


## What middleware is in Go

A middleware in Go is any `Handler` whose `ServeHTTP` calls the `ServeHTTP` of another `Handler`. For example we could log requests

```go
simpleLog := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s - %s\n", r.Method, r.URL.String()) //Log the request
    ducksHandler.ServeHTTP(w, r) //Now pass the request to ducksHandler
})

http.Handle("/simplelog", simpleLog)
```

So `simpleLog` is a middleware that logs a request's method and URL and then has `ducksHandler` handle the request. In Express a middleware is created by giving a handler function a `next` parameter.

```javascript
var simpleLog = function(req, res, next){
  console.log(req.method, '-', req.url)
  next()
}
app.all('/simplelog', simpleLog, ducksHandler)
```

So `ServeHTTP` is like `next` in Go.

A `ServeMux` is itself a middleware. Its `ServeHTTP` method looks through all of the `ServeMux`'s routes for a route that the request's URL matches and then calls the `ServeHTTP` of that route's `Handler`.

## Middleware chaining in Go

If you make a function that takes in a `Handler` and returns a `Handler` that runs some middleware functionality and then calls that `Handler`'s `ServeHTTP`, **you can pass in _any `Handler`_ to chain the middleware with that `Handler`**. For example, we could make `simpleLog` into a middleware chaining function like this:

```go
logRequest := func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s - %s\n", r.Method, r.URL.String())
        next.ServeHTTP(w, r)
    })
}

logRequestAndRaiseDuckVenomAwarenessChain := logRequest(ducksHandler)
http.Handle("/duckChain", logRequestAndRaiseDuckVenomAwarenessChain)
```

And since a middleware chaining function returns a `Handler`, that `Handler` can in turn be chained with other middlewares. For example, our duck awareness app could have an API route that does something like:

```go
http.Handle("/api/", logRequest(authenticate(serveData)))
```

For comparison, in Express that would look like:

```javascript
app.use('/api',      logRequest,authenticate,serveData)
```

`http.StripPrefix` also works like this as a middleware chaining function. It takes in a prefix string and a `Handler` and returns a middleware chain that removes that prefix from the request's URL and then gives the request to the `Handler` that was passed in.

## Making a middleware chain for all routes on a router

Routers like net/http `ServeMux`es and Gorilla mux `Router`s are themselves `Handler`s, (and as mentioned before, function as middlewares), so they can be added to middleware chains.

```go
mux := http.NewServeMux()

//...add some routes...

logAndServe := logRequest(mux)

server := &http.Server{
    Addr:    ":1123",
    Handler: logAndServe,
}
```

By chaining the actual router with a middleware chain and then **using the full middleware chain as the `Server`'s `Handler`**, all requests to the `Server` are processed by the middleware chain **and then are routed with the router.**

For comparison, in Express you would apply a middleware to all routes by adding a catch-all Express middleware route before all route-specific handlers:

```javascript
app.use(logRequest)

/* *** Other routes below *** */
```

## Middleware chaining libraries

Making a middleware chain in Go with one middleware chaining function calling another calling another can get to be ugly and cumbersome with all the parentheses. Luckily, there are several Go libraries for middleware chaining such as **Alice**, **MuxChain**, and **Negroni** that simplify making a middleware stack.

For this tutoial we will look at chaining with Alice and Negroni.

### Alice

Alice is a very simple middleware chaining package in Go that works by taking in a series of middleware chaining functions and a `Handler` and returning the middleware chain as one `Handler`. Its syntax is like this:

```go
//Create a chain of middleware constructors
middlewareChain :=
    alice.New(aMiddlewareConstructor, anotherMiddlewareConstructor)

//Create the complete chained middleware by passing a Handler into the
//chain's Then method; that Handler will be the last Handler of the
//middleware chain.
completeChain := middlewareChain.Then(aHandler)
```

Any function with the signature `func(http.Handler) http.Handler` can be used as a function in an Alice middleware chain.

So instead of:

```go
http.Handle("/api/", logRequest(authenticate(serveData)))
```

you could do:

```go
http.Handle("/api/", alice.New(logRequest, authenticate).Then(serveData))
```

Alice middleware chains have a syntax that feels similar to Express

```javascript
app.use('/api',                logRequest, authenticate,      serveData)
```

### Negroni

Negroni has some extra flexibility by defining its own data structure for its middleware stacks that can be either passed in middleware as regular `net/http` handler functions

```go
simpleLog := func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s - %s\n", r.Method, r.URL.String())
}

stack := negroni.New() //Create a Negroni middleware stack with negroni.New

//You can add an HTTP handler function to the middleware stack with
//UseHandlerFunc
stack.UseHandlerFunc(simpleLog)
```

or `Handler`s.

```go
//Add a handler like the router to the middleware stack with UseHandler
stack.UseHandler(serveMux)
```

Negroni also has its own `Handler` and handler function type you can use whose `ServeHTTP` method has the function signature:

`func (rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)`

Which looks a lot like the middleware functions in Express:

`function(req, res, next)`

A Negroni middleware stack is another type that implements `net/http` `Handler` interface, so if you have a Negroni middleware stack that includes a router in the stack, it can be used as the `Handler` for a `Server`

```go
server := &http.Server{
    Addr:    "1123",
    Handler: stack,
}
```
