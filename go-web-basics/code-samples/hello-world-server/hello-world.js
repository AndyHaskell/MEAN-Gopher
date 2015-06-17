var http = require('http')

function serveHello(req, res){
  var hello = 'Hello world!'
  res.writeHead(200, {
    'Content-Length' : hello.length,
    'Content-Type'   : 'text/plain'
  });
  res.end(hello, 'UTF-8')
}

http.createServer(serveHello).listen(1123)
