# Logging requests with Gorilla handlers

For logging HTTP requests, Gorilla has logging middleware in its `handlers` package.

To get the package run `go get github.com/gorilla/handlers`.

## Logging to stdout in Express with Morgan

```javascript
var app = express()

app.use(morgan('common'))

app.get('*', function(req, res){
  res.send('Sloths rule!')
})

http.createServer(app).listen(1123)
```

On our Express server, Morgan is the first middleware all requests go to. Each request is logged and then goes on to the other Express middleware functions for the routes on the server.

## Logging to stdout in Go with LoggingHandler

`LoggingHandler` is a function in the Gorilla handlers package that takes in an `io.Writer` and an `http.Handler` and constructs a new `Handler` chained with the `Handler` passed in.

This new `Handler`'s `ServeHTTP` logs the request and calls the `ServeHTTP` method of the `Handler` that was passed in.

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This is the catch-all route")
})

//As our io.Writer, we use os.Stdout and we chain our LoggingHandler
//with our ServeMux.
logAndServe := handlers.LoggingHandler(os.Stdout, mux)

//Now just pass in logAndServe as our server's Handler
server := &http.Server{
    Addr:    ":1123",
    Handler: logAndServe,
}
server.ListenAndServe()
```

`os.Stdout` is an `io.Writer`, so requests can be logged to standard output with that as our `LoggingHandler`'s writer.

On our `Server`, all of our requests first go to `logAndServe`, where the request is logged to standard output. Then, `logAndServe` passes the request to our `ServeMux` to route the request and serve a response like it normally would.



