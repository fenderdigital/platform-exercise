var userService = require("../services/userService");
var models = require("../models");

exports.all = function (request, response) {
    models.user.findAll().then(function (users) {
        response.json(users);
    })
};

exports.get = function (request, response) {
    "use strict";
    let id = request.params.id;
    userService.get(id).then(function (user) {
        response.json(user);
    }).catch(function () {
        response.status(404);
        return response.json({"message": "User not found"});
    })

};

exports.add = function (request, response) {
    "use strict";
    let id = request.params.id;
};

exports.delete = function (request, response) {
    "use strict";
    let id = request.params.id;
    userService.delete(id).then(function (value) {
        response.json(value);
    }).catch(function () {
        response.status(404);
        return response.json({"message": "User not found"});
    })
};

exports.set = function (request, response) {
    "use strict";
    let id = request.params.id;
    let user = exports.buildUserObject(request.body);

    userService.set(id, user).then(function (value) {
        response.json(value);
    });

};

exports.buildUserObject = function (body) {
    "use strict";
    let user = {};

    user.email = (body.email) ? body.email : null;
    user.name = (body.name) ? body.name : null;
    user.password = (body.password) ? body.password : null;

    return user;
};

exports.authenticate = function (request, response) {
    "use strict";
    let email = request.body.email;
    let password = request.body.password;
    userService.authenticate(email, password).then(function (value) {
        response.json(value);
    })
};

exports.logout = function (request, response) {
    "use strict";
    let token = request.body.token || request.query.token || request.headers['x-access-token'];

    response.json(userService.logout(token));
};
