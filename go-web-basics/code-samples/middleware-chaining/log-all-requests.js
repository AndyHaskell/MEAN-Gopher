var http    = require('http'),
    express = require('express')

var app = express()

var logRequest = function(req, res, next){
  console.log(req.method, '-', req.url)
  next()
}

var ducksHandler = function(req, res){
  res.send('Beware of ducks! Duck venom can turn people into ducks!')
}

app.use(logRequest)
app.use('/images', express.static(__dirname+'/public/images'))
app.use('/ducks', function(req, res){
  res.sendFile(__dirname+'/pages/ducks.html')
})
app.use('/', ducksHandler)

http.createServer(app).listen(1123)
