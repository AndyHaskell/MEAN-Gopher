var http       = require('http'),
    express    = require('express'),
    bodyParser = require('body-parser'),
    htmlEscape = require('html-escape')

var app = express()

var yourMethod = function(req, res){
  var reqMethod = req.method
  res.send('Your request method is: ' + reqMethod)
}

var serveOrderForm = function(req, res){
  res.sendFile(__dirname+'/pages/order-form.html')
}
var serveCoffeeShopOrderForm = function(req, res){
  res.sendFile(__dirname+'/pages/coffee-shop-order-form.html')
}

var serveSendOrder = function(req, res){
  var beverage = htmlEscape(req.body.beverage),
      name     = htmlEscape(req.body.name)

  res.send('One ' + beverage + ' coming right up, ' + name + '!')
}



//This only takes GET requests
app.get('/order-form', serveOrderForm)

//This route only takes POST requests and uses body-parser
//to get the POST data from the order form
app.post('/send-order', bodyParser.urlencoded({extended:true}),
                        serveSendOrder)

app.route('/coffee-shop')
   .get(serveCoffeeShopOrderForm)
   .post(bodyParser.urlencoded({extended:true}), serveSendOrder)

app.use('/', yourMethod)

http.createServer(app).listen(1123)
