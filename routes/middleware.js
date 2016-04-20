var env = process.env.NODE_ENV || "development";
var redisClient = require('../config/redis-database').redisClient;
var config = require("../config/config")[env];
var jwt = require('jsonwebtoken');
var Promise = require("bluebird");
var client = Promise.promisifyAll(redisClient);
var jwtVerifyAsync = Promise.promisify(jwt.verify, jwt);

exports.protected = function (req, res, next) {
    var token = req.body.token || req.query.token || req.headers['x-access-token'];

    if (token) {
        var mPromises = [client.getAsync(token), jwtVerifyAsync(token, config.secret)];
        Promise.all(mPromises).then(function (response) {
            var redis = response[0];
            if (redis) {
                req.decoded = response[1];
                next();
            } else {
                return res.status(403).send({
                    success: false,
                    message: 'No active session'
                });
            }
        }).catch(function (e) {
            return res.json(e);
        });
    } else {
        return res.status(403).send({
            success: false,
            message: 'No token provided.'
        });
    }
};
