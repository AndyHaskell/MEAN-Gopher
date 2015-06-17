# Static file serving
## Making a file server route
### In Express:

```javascript
var app = express()
app.use('/img', express.static(__dirname+'/public/images'))
```
### In Go:
```go
mux := http.NewServeMux()
imgServer := http.StripPrefix("/img/",
    http.FileServer(http.Dir("public/images")))
mux.Handle("/img/", imgServer)
```

`http.Dir` accesses a directory, in our case `public/images`.

`http.FileServer` takes in a directory and creates a `Handler` that serves requests by looking for files in that directory that match the request.

`http.StripPrefix` takes in a prefix and a `Handler` and constructs our final `Handler` for the route. The `Handler` from `StripPrefix` is a middleware that removes the prefix it was given from the request's URL and then calls the `ServeHTTP` method of the `Handler` it was given, now with the request URL's prefix removed.

- So a request to `/img/sloth.jpg` would get the prefix `/img/` removed. Then the `FileServer`'s `ServeHTTP` would be called with the request URL now being `sloth.jpg`.
- The `FileServer` would then look for the image `public/images/sloth.jpg`
- Without `StripPrefix`, the request the `FileServer` got would have the URL **`img/sloth.jpg`**, so the `FileServer` would look in **`public/images/img/sloth.jpg`**.
## A convenient work-around for Go file servers
For a more convenient syntax for setting up file server routes, you can use this function:
```go
func FileServerRoute (mux *http.ServeMux, path, dir string) {
    mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}
```
For example:
```go
mux := http.NewServeMux()
FileServerRoute(mux, "/img/", "public/images")
```
## Serving static HTML files
For serving individual files, much like how in Express there's a `sendFile` response method, `net/http` in Go has a built-in `ServeFile` method.
### In Express
```javascript
app.use(function(req, res){
  res.sendFile(__dirname+'/pages/index.html')
})
```
### In Go
```go
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "pages/index.html")
})
```
