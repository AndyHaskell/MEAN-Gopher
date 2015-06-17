# Starting a basic server

## The Hello world example

### In Node:

```javascript
var http = require('http')

function serveHello(req, res){
  var hello = 'Hello world!'
  res.writeHead(200, {
    'Content-Length' : hello.length,
    'Content-Type'   : 'text/plain'
  });
  res.end(hello, 'UTF-8')
}

http.createServer(serveHello).listen(1123)
```

### In Go:

```go
package main

import (
    "fmt"
    "net/http"
)

func serveHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    http.HandleFunc("/", serveHello)
    http.ListenAndServe(":1123", nil)
}
```

Since `serveHello` in Node is passed to `http.createServer`, that function handles all requests to the Node server.

Likewise, since `serveHello` in Go matches the `"/"` route, all requests to the Go server are handled by that route where the requests are handled by `serveHello`.

## Making a server as its own object

### In Node:

```javascript
var server = http.createServer(serveHello)
```

### In Go:

```go
server := &http.Server{
    Addr: ":1123",
    Handler: http.DefaultServeMux,
}
```

`Addr` is what port our server will listen for requests on (in our case Port 1123).

`Handler` is a Go interface for handling requests to the Server. In our case, we're using `DefaultServeMux`, which is the Handler where routes get registered to when you call `http.Handle` or `HandleFunc`.

There are many other parameters we can give the server as well, which are documented in [net/http's documentation](http://golang.org/pkg/net/http/).


## Starting a server from a Server object

### In Node:
```javascript
server.listen(1123)
```

### In Go:
```go
server.ListenAndServe()
```

Behind the scenes, `http.ListenAndServe()` creates a `Server` and then **calls its `ListenAndServe` method**.