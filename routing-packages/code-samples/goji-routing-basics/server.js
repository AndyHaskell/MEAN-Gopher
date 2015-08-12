var http    = require('http'),
    express = require('express')

var app = express();

//Make a plain path
app.use('/sloths', function(req, res){
  res.send('Sloths rule!')
})

//Middleware chaining with request contexts
var youreNo1000000 = function(req, res, next){
  req.hitNumber = 1000000
  next()
}
var serveHitNumber = function(req, res){
  res.send('You are totally viewer number ' + req.hitNumber + '!')
}

//Route parameters
app.use('/:flavor/tea', function(req, res){
  flavor = req.params.flavor
  res.send('I could go for some ' + flavor + ' tea!')
})

//Static file server
app.use('/img', express.static(__dirname+'/public/images'))

//GET-specific route
app.get('/get-route', function(req, res){
  res.send('This route only responds to GET requests')
})

//Regular expression matching with a RegExp
app.get(/\/regexp\/(coffee)+$/, function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})

//Regular expression matching with a pattern
app.get('(/pattern/)(coffee)+', function(req, res){
  res.send('Lemurs = sloths that had too much coffee')
})

app.use('/', youreNo1000000, serveHitNumber)

http.createServer(app).listen(1123)
