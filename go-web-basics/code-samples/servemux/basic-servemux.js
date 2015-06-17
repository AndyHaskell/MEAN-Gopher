var http    = require('http'),
    express = require('express')

var app = express()

app.use('*', function(req, res){
  res.send('Sloths rule!')
})

var server = http.createServer(app)
server.listen(1123)
