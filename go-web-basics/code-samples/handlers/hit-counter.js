var http    = require('http'),
    express = require('express')

var app = express()
var counter = 0

//Logger middleware, run by all requests
app.use(function(req, res, next){
  console.log('Request to', req.url)
  next()
})

//Serve the /sloths route
app.use('/sloths', function(req, res){
  res.send('Sloths rule!')
})

//Catch-all route, serves the hit counter
app.use(function(req, res) {
  counter++
  res.send("This app's hit count: " + counter)
})

http.createServer(app).listen(1123)
