# Routing basics in Gorilla mux

Gorilla mux is a Go package used for creating an HTTP router that's very similar to the `ServeMux` router in `net/http` but with additional path-matching capabilities built-in. It can match plain paths, paths with parameters, regular expressions, and even custom routing functions for complex routes.

The syntax for using a Gorilla mux `Router` is almost exactly the same as for using a regular `ServeMux`, so it's kind of like Express routing for Go.

You can get it by running `go get github.com/gorilla/mux`.

## Starting a router and adding a plain path:
### In Express:
```javascript
var app = express();
app.get('/sloths', function(req, res){
  res.send('Sloths rule!')
})
http.createServer(app).listen(1123)
```
### In Gorilla:
To create a Gorilla mux router, you use `mux.NewRouter`. Like with `net/http`, you can register a route to a `Handler` or handler function with `Router.Handle` and `Router.HandleFunc`.
```go
m := mux.NewRouter()
m.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Sloths rule!")
})

server := &http.Server {
    Addr:    ":1123",
    Handler: m,
}
server.ListenAndServe()
```
Note that a Gorilla mux `Router` has a `ServeHTTP` method and is therefore a `Handler`, so it can be used as a `Server`'s handler the same way a `ServeMux` can.

## Path prefixes
### An image server route with a net/http ServeMux:
```go
serveMux.Handle("/img/", http.StripPrefix("/img/",
	http.FileServer(http.Dir("public/images"))))
```

### With a Gorilla mux Router:
Unlike in a `ServeMux`, in a Gorilla mux router, plain paths ending in a slash **only match themselves**, so to use a path as a prefix, you instead use `Router.PathPrefix`.
```
m.PathPrefix("/img/").Handler(
    http.StripPrefix("/img/", http.FileServer(http.Dir("public/images"))))
```

`Router.PathPrefix` makes a new Gorilla mux `Route`, which you can then give a Handler to with `Route.Handler` or a HandlerFunc with `Route.HandlerFunc`.

Because of this Gorilla routing rule, you would also do `PathPrefix` for a catch-all route in Gorilla.

```go
m.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This route matches all requests.")
})
```

## Route parameters
### In Express:
```javascript
app.get('/:flavor/tea', function(req, res){
  flavor = req.params.flavor
  res.send('I could go for some ' + flavor + ' tea!')
})
```

### In Gorilla:

In Gorilla you use a curly braces `{paramName}` format for route parameters and you fetch a request's route parameters by calling `mux.Vars` with the `*Request`.
```go
m.HandleFunc("/{flavor}/tea", func(w http.ResponseWriter, r *http.Request){
    routeParams := mux.Vars(r)
    fmt.Fprintf(w, "I could go for some %q tea!", routeParams["flavor"])
})
```

## Regular expressions
### In Express:
```javascript
//With a RegExp
app.get(/\/regexp\/(coffee)+$/, function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})

//With a pattern
app.get('(/pattern/)(coffee)+', function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})
```

### In Gorilla:

In Gorilla, you add regular expression route parameters in the format `{parameterName:regularExpression}`. 
```go
m.HandleFunc(`/{drink:(coffee)+}`,
    func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Lemurs = sloths that had too much coffee")
})
```
To make it easier to work with regex escape characters in regular expressions in Go, I recommend using backtick-quoted strings to define paths that use regular expression matching in Gorilla mux.
