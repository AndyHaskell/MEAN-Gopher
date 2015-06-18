var http    = require('http'),
    express = require('express')

var app = express();

//Plain route
app.get('/sloths', function(req, res){
  res.send('Sloths rule!')
})

//Route params
app.get('/:flavor/tea', function(req, res){
  flavor = req.params.flavor
  res.send('I could go for some ' + flavor + ' tea!')
})

//Regular expression matching with a RegExp
app.get(/\/regexp\/(coffee)+$/, function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})

//Regular expression matching with a pattern
app.get('(/pattern/)(coffee)+', function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})

//Catch-all route
app.get('*', function(req, res){
  res.send('This route matches all requests.')
})

http.createServer(app).listen(1123)
