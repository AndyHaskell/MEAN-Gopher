var http    = require('http'),
    express = require('express')

var app = express()

var ducksHandler = function(req, res){
  res.send('Beware of ducks! Duck venom can turn people into ducks!')
}

var simpleLog = function(req, res, next){
  console.log(req.method, '-', req.url)
  next()
}

app.all('/simplelog', simpleLog, ducksHandler)
app.use('/', ducksHandler)

http.createServer(app).listen(1123)
