var http    = require('http'),
    express = require('express'),
	morgan  = require('morgan')

var app = express()

app.use(morgan('common'))

app.get('*', function(req, res){
  res.send('Sloths rule!')
})

http.createServer(app).listen(1123)
