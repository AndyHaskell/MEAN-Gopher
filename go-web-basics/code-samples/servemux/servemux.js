var http    = require('http'),
    express = require('express')

var app = express()

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

var server = http.createServer(app)
server.listen(1123)
