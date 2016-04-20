var env = process.env.NODE_ENV || "development";
var config = require("./config")[env];
var redis = require('redis');
var redisClient = redis.createClient(config.redis, {no_ready_check: true});

redisClient.on('error', function (err) {
    console.log('Error ' + err);
});

redisClient.on('connect', function () {
    console.log('Redis is ready');
});

exports.redis = redis;
exports.redisClient = redisClient;
