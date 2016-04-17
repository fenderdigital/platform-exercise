var env = process.env.NODE_ENV || "development";
var config = require("../config/config")[env];
var redisClient = require('../config/redis-database').redisClient;
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
            if (result) {
                var token = exports.createToken(mUser);
                redisClient.set(token, JSON.stringify(mUser));
                redisClient.expire(token, config.expiration);
                value.token = token;
            }
            value.success = result;
            value.message = (result) ? SUCCESS : FAIL;

            return resolve(value);
        }).catch(function (e) {
            return reject(e);
        });
    });
};

exports.logout = function (token) {
    "use strict";
    redisClient.del(token);

    return {success: true, message: "logged out"}
};

exports.createToken = function (user) {
    "use strict";
    let tokenUser = {
        name: user.name,
        email: user.email
    };

    return jwt.sign(tokenUser, config.secret, {
        expiresIn: config.expiration
    });

};
