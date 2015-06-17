var http    = require('http'),
    express = require('express')

var app = express()

app.use('/img', express.static(__dirname+'/public/images'))

app.use(function(req, res){
  res.sendFile(__dirname+'/pages/index.html')
})

http.createServer(app).listen(1123)
