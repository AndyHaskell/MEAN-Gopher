# Routing in Go with a ServeMux
## Router data types
### In JavaScript (with an Express app as the router)
In Express, an Express app is a router and it can be used as the request listener for a `Server` in `http`. 
```javascript
var app = express()

app.use('*', function(req, res){
  res.send('Sloths rule!')
})

var server = http.createServer(app)
server.listen(1123)
```
### In Go (with a ServeMux as a router)
`net/http` has the built-in router `http.ServeMux` as a basic router with basic path matching, which can be used as the Handler for a `Server` in `net/http`.
```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Sloths rule!")
})
    
server := &http.Server {
    Addr: ":1123",
    Handler: mux,
}
server.ListenAndServe()
```
## Go path matching
Handlers in a `ServeMux` mapped to a path that doesn't end with a slash only match that exact path. Handlers mapped to a path that does end with a slash match that exact path, and optionally, a slash and everything after it.
```go
//Matches ONLY /sloths, not /sloths/ or /sloths/are-awesome
mux.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Sloths rule!")
})

//Matches /kangaroos, /kangaroos/, and /kangaroos/tree-kangaroos
mux.HandleFunc("/kangaroos/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Kangaroos for the win!")
})

//Since "/" ends with a slash, it matches all URL paths, so "/"
//is the catch-all route.
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Lemurs get the catch-all route! Lemurs are where it's at!")
})
```
### Those routes in Express
```javascript
//Matches ONLY /sloths and /sloths/, not /sloths/are-awesome
app.get('/sloths', function(req, res){
  res.send('Sloths rule!')
})

//Matches /kangaroos, /kangaroos/, and /kangaroos/tree-kangaroos
app.get('/kangaroos(/*)?', function(req, res){
  res.send('Kangaroos for the win!')
})

//Since "*" matches all URL paths, so "*" is the catch-all route.
app.get('*', function(req, res){
  res.send("Lemurs get the catch-all route! Lemurs are where it's at!")
})
```
## Routing order
Unlike in Express, where a request is matched to the earliest-defined route its URL matches, in a `ServeMux` a request is matched to whichever route it matches that has the longest path.
### In Express
```javascript
//tea/hibiscus matches this route
app.get('/tea(/*)?', function(req, res){
  res.send('One tea coming right up!')
})

//Since tea/hibiscus already matched the above route, this route
//isn't matched, so this route should be defined first if we want this
//route to be matched first.
app.get('/tea/hibiscus', function(req, res){
  res.send('One hibiscus tea coming right up!')
})
```
### In Go
```go
// tea/hibiscus matches this route, but the route below is longer
mux.HandleFunc("/tea/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "One tea coming right up!")
})

//This route has a longer path so a request to tea/hibiscus matches this
//route, not the one above.
mux.HandleFunc("/tea/hibiscus", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "One hibiscus tea coming right up!")
})
```

## DefaultServeMux

`net/http` has a built-in `ServeMux` called `http.DefaultServeMux`.

When you call `http.Handle` or `http.HandleFunc`, the route you create is mapped to `DefaultServeMux`.

```go
http.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Sloths rule!")
})

//is equivalent to

http.DefaultServeMux.HandleFunc("/sloths",
    func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Sloths rule!")
})
```

Behind the scenes, if a `Server`'s `Handler` property is `nil`, it will handle requests with `DefaultServeMux`.
