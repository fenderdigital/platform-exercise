var env = process.env.NODE_ENV || "development";
var config = require("../config/config")[env];
var jwt = require('jsonwebtoken');
var Promise = require("bluebird");

var jwtVerifyAsync = Promise.promisify(jwt.verify, jwt);

exports.protected = function (req, res, next) {
    var token = req.body.token || req.query.token || req.headers['x-access-token'];

    if (token) {
        jwtVerifyAsync(token, config.secret).then(function (decoded) {
            req.decoded = decoded;
            next();
        }).catch(function (e) {
            return res.json({success: false, message: 'Failed to authenticate token.'});
        });
    } else {
        return res.status(403).send({
            success: false,
            message: 'No token provided.'
        });
    }
};
