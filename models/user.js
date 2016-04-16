"use strict";
module.exports = function (sequelize, DataTypes) {
    return sequelize.define('user', {
        name: DataTypes.STRING,
        email: {
            type: DataTypes.STRING,
            unique: true,
            validate: {
                isEmail: true
            }
        },
        password: {
            type: DataTypes.STRING,
            validate: {
                is: {
                    args: "^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$",
                    msg: "Minimum 8 characters at least 1 Alphabet, 1 Number and 1 Special Character"
                }
            }
        }
    });
};
