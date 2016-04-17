var env = process.env.NODE_ENV || "development";
var port = process.env.PORT || 8080;

var config = require("./config/config")[env];
var express = require('express');
var app = express();
var bodyParser = require('body-parser');
var morgan = require('morgan');
var models = require("./models");
var tokenManager = require('./config/redis-database');
var Promise = require("bluebird");
var bcrypt = Promise.promisifyAll(require('bcrypt'));


var routes = require('./routes/index');
var users = require('./routes/user');

app.set('superSecret', config.secret);

app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());
app.use(morgan('dev'));

/*
 app.use(function(err, req, res, next) {
 res.status(err.status || 500);
 res.render('error', {
 message: err.message,
 error: (app.get('env') === 'development') ? err : {}
 });
 });
 */

models.sequelize.sync().then(function () {
    return bcrypt.hashAsync("12345678", config.saltRounds);
}).then(function (password) {
     return models.user.create({
     name: "Alejandro",
     email: "arangel@fender.com",
     password: password
     });

});


app.use('/', routes);
app.use('/users', users);


app.listen(port);
console.log('Magic happens at http://localhost:' + port);

