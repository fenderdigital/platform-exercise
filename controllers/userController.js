var userService = require("../services/userService");
var models = require("../models");

exports.all = function (request, response) {
    models.user.findAll().then(function (users) {
        response.json(users);
    })
};

exports.get = function (request, response) {

};

exports.add = function (request, response) {

};

exports.set = function (request, response) {

};

exports.authenticate = function (request, response) {
    "use strict";
    let email = request.body.email;
    let password = request.body.password;
    userService.authenticate(email, password).then(function (value) {
        response.json(value);
    })
};
