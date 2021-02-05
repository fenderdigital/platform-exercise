
const constants = require('../constants')
var jsonwebtoken = require("jsonwebtoken");
const db = require("../models");
const config = require("../configurations/auth.config");
const User = db.user;

  exports.allAccess = (req, res) => {
    res.status(200).send({"message": constants.DISPLAY_MESSAGE});
  };
  exports.userPrivate = (req, res) => {
    res.status(200).send(constants.USER_PRIVATE);
  };
    
  exports.nameUpdate = (req, res, next) => {
    let token = req.headers["x-access-token"];

    jsonwebtoken.verify(token, config.secret, (err, decoded) => {
      if (err) {
        return res.status(401).send({
          message: constants.UNAUTH_FORBIDDEN
        });
      }

      User.update(
        {name: req.body.name},
        {returning: true, where: {id: decoded.id} }
      )
      .then(function([ rowsUpdate, [updatedUser] ]) {
        res.json(updatedUser)
      })
      .catch(next)
   
    }); 
  }