var env = process.env.NODE_ENV || "development";
var config = require("../config/config")[env];

var models = require("../models");
var Promise = require("bluebird");
var jwt = require('jsonwebtoken');


var bcrypt = Promise.promisifyAll(require('bcrypt'));


var SUCCESS = "Success login";
var FAIL = 'Authentication failed. Wrong password.';

exports.authenticate = function (mail, password) {
    "use strict";
    var value = {};
    var mUser;
    return new Promise(function (resolve, reject) {
        models.user.findOne({
            where: {
                email: mail
            }
        }).then(function (user) {
            mUser = user;
            return bcrypt.compareAsync(password, user.password)
        }).then(function (result) {
            console.log(result);
            if (result) {
                value.token = createToken(mUser);
            }
            value.success = result;
            value.message = (result) ? SUCCESS : FAIL;

            return resolve(value);
        }).catch(function (e) {
            return reject(e);
        });
    });
};

createToken = function (user) {
    "use strict";
    let tokenUser = {
        name: user.name,
        email: user.email
    };

    return jwt.sign(tokenUser, config.secret, {
        expiresIn: 86400 // expires in 24 hours
    });

};
