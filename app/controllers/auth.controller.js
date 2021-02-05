const db = require("../models");
const config = require("../configurations/auth.config");
const User = db.user;
var bcrypt = require("bcryptjs");
var jsonwebtoken = require("jsonwebtoken");
const constants = require('../constants')

exports.signup = (req, res) => {

  User.create({
    name: req.body.name,
    email: req.body.email,
    password: bcrypt.hashSync(req.body.password, 8)
  })
    .then(user => {
      res.send({ message: constants.REGISTERED_SUCCESS });
    })
    .catch(err => {
      res.status(500).send({ message: err.message });
    });
};
exports.signin = (req, res) => {
  User.findOne({
    where: {
      email: req.body.email
    }
  })
    .then(user => {
      if (!user) {
        return res.status(404).send({ message: constants.USER_NOT_FOUND});
      }
      var passwordIsValid = bcrypt.compareSync(
        req.body.password,
        user.password
      );

      if (!passwordIsValid) {
        return res.status(401).send({
          accessToken: null,
          message: constants.INVALID_PASSWORD
        });
      }

      var token = jsonwebtoken.sign({ id: user.id }, config.secret, {
        expiresIn: 86400 
      });

      res.status(200).send({
        id: user.id,
        name: user.name,
        email: user.email,
        accessToken: token
      });

    })
    .catch(err => {
      res.status(500).send({ message: err.message });
    });
};