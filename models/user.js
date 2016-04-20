"use strict";
module.exports = function (sequelize, DataTypes) {
    return sequelize.define('user', {
        name: {
            type: DataTypes.STRING,
            notNull: true
        },
        email: {
            type: DataTypes.STRING,
            unique: true,
            validate: {
                isEmail: true
            },
            notNull: true
        },
        password: {
            type: DataTypes.STRING,
            notNull: true
        }
    });
};
