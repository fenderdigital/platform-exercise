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

exports.add = function (mUser) {
    return new Promise(function (resolve, reject) {
        models.user.create({
            name: mUser.name,
            email: mUser.email,
            password: mUser.password
        }).then(function (value) {
            return resolve(value);
        }).catch(function (e) {
            return reject(e);
        });
    })
};

exports.set = function (id, mUser) {
    return new Promise(function (resolve, reject) {
        models.user.findById(id).then(function (user) {
            if (user) {
                return exports.updateValues(user, mUser)
            }
            return reject();
        }).then(function (updatedUser) {
            return updatedUser.save()
        }).then(function (value) {
            return resolve(value)
        }).catch(function (e) {
            return reject(e);
        })
    });
};

exports.delete = function (id) {
    return new Promise(function (resolve, reject) {
        models.user.findById(id).then(function (mUser) {
            if (mUser) {
                return mUser.destroy();
            }
            return reject();
        }).then(function () {
            return resolve({success: true});
        }).catch(function (e) {
            console.log(e);

            return reject(e)
        })
    });

};

exports.updateValues = function (user, mUser) {
    return new Promise(function (resolve, reject) {
        user.email = (mUser.email) ? mUser.email : user.email;
        user.name = (mUser.name) ? mUser.name : user.name;
        if (mUser.password) {
            bcrypt.hashAsync(mUser.password, config.saltRounds).then(function (passwordHash) {
                user.password = passwordHash;
                return resolve(user);
            });
        }
        else
            return resolve(user);

    });
};

exports.get = function (id) {
    return new Promise(function (resolve, reject) {
        models.user.findById(id).then(function (user) {
            if (user) {
                return resolve(user);
            }

            return reject();
        }).catch(function (e) {
            return reject(e);
        })
    });
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
