"use strict";

var fs = require("fs");
var path = require("path");
var Sequelize = require("sequelize");
var env = process.env.NODE_ENV || "development";
var config = require("../config/config")[env];
var sequelize = new Sequelize(config.url,
    {
        dialect: config.dialect,
        dialectOptions: {
            ssl: true
        }
    });

var database = {};

fs.readdirSync(__dirname)
    .filter(function (file) {
        return (file.indexOf(".") !== 0) && (file !== "index.js");
    })
    .forEach(function (file) {
        var model = sequelize.import(path.join(__dirname, file));
        database[model.name] = model;
    });

Object.keys(database).forEach(function (modelName) {
    if ("associate" in database[modelName]) {
        database[modelName].associate(database);
    }
});

database.sequelize = sequelize;
database.Sequelize = Sequelize;

module.exports = database;
