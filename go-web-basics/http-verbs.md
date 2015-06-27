# HTTP verb basics in Go

## Getting a request's HTTP method

### In Express

```javascript
var app = express()
var yourMethod = function(req, res){
  var reqMethod = req.method
  res.send('Your request method is: ' + reqMethod)
}
app.use('/', yourMethod)
```

### In Go net/http

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    reqMethod := r.Method
    if reqMethod == "" {
        reqMethod = "GET"
    }
    
    fmt.Fprintf(w, "Your request method is %s", reqMethod)
})
```

In both Node `http` and Go `net/http`, requests have a method property. In Go, if a request's Method is blank, that means it's a `GET` request.

## Making routes specific to an HTTP verb

### In Express

```javascript
var serveOrderForm = function(req, res){
  res.sendFile(__dirname+'/pages/order-form.html')
}

var serveSendOrder = function(req, res){
  var beverage = htmlEscape(req.body.beverage),
      name     = htmlEscape(req.body.name)

  res.send('One ' + beverage + ' coming right up, ' + name + '!')
}

//This only takes GET requests
app.get('/order-form', serveOrderForm)

//This route only takes POST requests and uses body-parser
//to get the POST data from the order form
app.post('/send-order', bodyParser.urlencoded({extended:true}),
                        serveSendOrder)
```

### In Go

```go
serveOrderForm := func(w http.ResponseWriter, r *http.Request) {
    //Restrict the route to only GET requests
    if (r.Method == "GET" || r.Method == "") {
        http.ServeFile(w, r, "pages/order-form.html")
    } else {
        http.Error(w, "405 Method Not Allowed", 405)
    }
}

serveSendOrder := func(w http.ResponseWriter, r *http.Request) {
    //Restrict the route to only POST requests
    if (r.Method == "POST") {
        //Parse the POST data with r.ParseForm() and
        //get the data with r.Form.Get()
        r.ParseForm()
        beverage := html.EscapeString(r.Form.Get("beverage"))
        name     := html.EscapeString(r.Form.Get("name"))

        fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
            beverage, name)
    } else {
        http.Error(w, "405 Method Not Allowed", 405)
    }
}

mux.HandleFunc("/order-form", serveOrderForm)
mux.HandleFunc("/send-order", serveSendOrder)
```

In Go's `net/http` package, you can restrict a route to specific HTTP verbs with an if statement checking the request's `Method`.

## Routes handling more than one verb

### In Express

```javascript
//This is just so we have an order form HTML file that will send
//POST data to /coffee-shop, not /send-order
var serveCoffeeShopOrderForm = function(req, res){
  res.sendFile(__dirname+'/pages/coffee-shop-order-form.html')
}

//You can handle multiple routes at once with route()
app.route('/coffee-shop')
   .get(serveCoffeeShopOrderForm)
   .post(bodyParser.urlencoded({extended:true}), serveSendOrder)
```

### In Go (with one handler function)

```go
mux.HandleFunc("/coffee-shop", func(w http.ResponseWriter, r *http.Request){
    if (r.Method == "GET" || r.Method == "") {
        http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
    } else if (r.Method == "POST") {
        r.ParseForm()
        beverage := html.EscapeString(r.Form.Get("beverage"))
        name     := html.EscapeString(r.Form.Get("name"))

        fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
            beverage, name)
    } else {
        http.Error(w, "405 Method Not Allowed", 405)
    }
})
```

You can just combine the handlers into one big handler function with a branch of the if statement for each HTTP method, but that's not modular.

### In Go (modularized into multiple handler functions)


```go
orderForm := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
})
sendOrder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    beverage := html.EscapeString(r.Form.Get("beverage"))
    name     := html.EscapeString(r.Form.Get("name"))

    fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
        beverage, name)
})

mux.HandleFunc("/coffee-shop", func(w http.ResponseWriter, r *http.Request){
    if (r.Method == "GET" || r.Method == "") {
        orderForm.ServeHTTP(w, r)
    } else if (r.Method == "POST") {
        sendOrder.ServeHTTP(w, r)
    } else {
        http.Error(w, "405 Method Not Allowed", 405)
    }
})
```

You can modularize a Go `net/http` handler function handling different HTTP verbs with different functionality by having a **central handler function that passes the request and response writer to a different `Handler` for each HTTP verb by calling that `Handler`'s `ServeHTTP` method!**

So if `/modular-cafe` gets a `GET` request, the request is handled with `orderForm` ans if it gets a `POST` request, it's handled with `sendOrder`, but if it gets a request with a different method like `PUT` or `DELETE`, we get a **405 Method Not Allowed** error.

That central handler function calling other `Handler`s' `ServeHTTP` methods is an example of a web middleware in Go that chains with either `orderForm` or `sendOrder`.

## Routing packages

In Go `net/http` at the time I am writing this, there isn't any built-in way to restrict a route to only certain HTTP methods. Luckily, if you want something more sleek like what you get Express, there are a ton of HTTP routing packages in the Go community. Here are a couple examples for the `/coffee-shop` route with Gorilla mux and Goji.

### Our handler functions

```go
serveOrderForm := func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "pages/order-form.html")
}
serveCoffeeShopOrderForm := func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "pages/coffee-shop-order-form.html")
}

serveSendOrder := func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    beverage := html.EscapeString(r.Form.Get("beverage"))
    name     := html.EscapeString(r.Form.Get("name"))

    fmt.Fprintf(w, "<body>One %s coming right up, %s!</body>",
        beverage, name)
}
```

### Serving our handler functions in Gorilla mux

```go
m := mux.NewRouter() //Create a Gorilla mux Router

//You restrict routes to specific HTTP methods with Route.Methods()
m.HandleFunc("/order-form", serveOrderForm).Methods("GET")
m.HandleFunc("/send-order", serveSendOrder).Methods("POST")

//You can also register multiple handlers to the same path in Gorilla and
//having Gorilla resolve which one to serve with Route.Methods
m.Path("/coffee-shop").HandlerFunc(serveCoffeeShopOrderForm).Methods("GET")
m.Path("/coffee-shop").HandlerFunc(serveSendOrder).Methods("POST")
```

### Serving our handler functions in Goji

```go
m := web.New() //Create a Goji Mux

//A Goji Mux comes with Get and Post methods for registering routes for
//specific HTTP methods
m.Get("/order-form", serveOrderForm)
m.Post("/send-order", serveSendOrder)

//You can register handlers to the same route with Get and Post
m.Get("/coffee-shop", serveCoffeeShopOrderForm)
m.Post("/coffee-shop", serveSendOrder)
```
