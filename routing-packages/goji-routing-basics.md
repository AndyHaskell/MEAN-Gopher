# Routing basics in Goji

Goji is a web microframework that gives you routing capabilities that feel a lot like those in a `net/http` `ServeMux`, but with path matching capabilities like those in Express.js or in [Gorilla mux](http://www.gorillatoolkit.org/pkg/mux), such as being able to match paths with parameters or regular expressions and making HTTP verb-specific routes.

Goji has fast routing and a familiar routing syntax for anyone coming from an Express.js background. For this tutorial, we will be taking the examples from [this Gorilla mux tutorial](https://github.com/AndyHaskell/MEAN-Gopher/blob/master/routing-packages/gorilla-mux-basics.md) and doing them in Goji.

You can get Goji by running `go get github.com/goji`.

## Starting a router and adding a plain path:
### In Express:
```javascript
var app = express();
app.use('/sloths', function(req, res){
  res.send('Sloths rule!')
})
http.createServer(app).listen(1123)
```
### In Goji:
```go
m := web.New()
m.Handle("/sloths", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Sloths rule!")
})

server := &http.Server {
    Addr:    ":1123",
    Handler: m,
}
server.ListenAndServe()
```
To create a Goji web `Mux`, you use `web.New`. Like with `net/http`, you can register a route to a `Handler` or handler function with the `Mux`'s `Handle` method. Unlike in `net/http` and Gorilla mux, however, **a Goji `Mux`'s `Handle` method can take in either a `Handler` or handler function, rather than having specific `Handle` and `HandleFunc` methods**

Note that a Goji `Mux` has a `ServeHTTP` method and is therefore a `net/http` `Handler`, so it can be used as a `Server`'s handler the same way a `ServeMux` or Gorilla `Router` can.

## Request contexts

Go `net/http` `Request`s aren't as flexible as Node.js requests, so Goji fixes this with a context type called `web.C`.

While a `net/http` `Handler` takes in just a `ResponseWriter` and `Request` for handling requests with its `ServeHTTP method`, a Goji `Handler` has a method `ServeHTTPC` with this signature:

```go
ServeHTTPC(C, http.ResponseWriter, *http.Request)
```

And Goji methods that can create routes in `net/http` `Handler`s can also take in Goji `Handler`s and handler functions (with the `ServeHTTPC` signature), giving you request context functionality if you need it.

### Adding request context values in Express

```javascript
var youreNo1000000 = function(req, res, next){
  req.hitNumber = 1000000
  next()
}
var serveHitNumber = function(req, res){
  res.send('You are totally viewer number ' + req.hitNumber + '!')
}

app.use('/', youreNo1000000, serveHitNumber)
```

### Adding request context values in Goji

```go
//A simple chain of Goji handler functions
func youreNo1000000(c web.C, w http.ResponseWriter, r *http.Request){
    c.Env["hitNumber"] = 1000000
    serveHitNumber(c, w, r)
}
func serveHitNumber(c web.C, w http.ResponseWriter, r *http.Request){
    hitNumber := c.Env["hitNumber"]
    fmt.Fprintf(w, "You're totally viewer number %d!", hitNumber)
}

func main(){
    m := web.New()
    
    //Initialize a request context's Env and make the route
    m.Use(middleware.EnvInit)
    m.Handle("/", youreNo1000000)
}
```

A Goji `C` stores variables in a map called `c.Env`. To use its `Env` **your Goji `Mux` must import `"github.com/zenazn/goji/web/middleware"` and have the middleware `middleware.EnvInit` in the `Mux`**.

## Route parameters
### In Express
```javascript
app.use('/:flavor/tea', function(req, res){
  flavor = req.params.flavor
  res.send('I could go for some ' + flavor + ' tea!')
})
```

### In Goji

```go
m.Handle("/:flavor/tea", func(c web.C, w http.ResponseWriter, r *http.Request){
    flavor := c.URLParams["flavor"]
    fmt.Fprintf(w, "I could go for some %s tea!", flavor)
})
```
In Goji, you use the same Sinatra-like :routeParameter syntax you would use in Express. To fetch the parameters from the route, a Goji `C` contains a `URLParams` map.

## Path prefixes
### An image server route in Express
```javascript
app.use('/img', express.static(__dirname+'/public/images'))
```

### In net/http
```go
serveMux.Handle("/img/", http.StripPrefix("/img/",
	http.FileServer(http.Dir("public/images"))))
```

### In Goji
In Goji, part of the Sinatra-like path matching syntax is using `*`s in paths. So `/img/*` matches paths with the prefix `/img/`.
```
m.Handle("/img/*", http.StripPrefix("/img/", 
    http.FileServer(http.Dir("public/images"))))
```

Because of this routing rule, `/*` can be used as a catch-all route.

```go
m.Handle("/*", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This route matches all requests.")
})
```

## HTTP verb-specific routes

A Goji `Mux`'s `Handle` method creates a route that handles requests that match the path regardless of their HTTP verb. Goji also offers syntax for making routes that are specific to to one HTTP verb, which look a lot like the methods you would use on an Express.js app:

### A GET route in Express
```javascript
app.get('/get-route', function(req, res){
  res.send('This route only responds to GET requests')
})
```

### In Goji
```go
m.Get("/get-route", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This route only responds to GET requests")
})
```

## Regular expressions
### In Express
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

### In Goji

```go
coffeeRegexp := regexp.MustCompile(`^/(coffee)+$`)
m.Get(coffeeRegexp, func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Lemurs = sloths that had too much coffee")
})
```

In Goji, in addition to routes in `Handle`, `Get`, `Post`, etc. being able to match paths with Sinatra-like syntax, routes' patterns can also be regular expressions. These regular expressions are `regexp.Regexp` objects.



