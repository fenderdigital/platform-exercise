/*
 * Middleware functions
*/
const db = require("../models");
const User = db.user;

checkDuplicateEmail = (req, res, next) => {

  User.findOne({
    where: {
      email: req.body.email
    }
  }).then(user => {
      if (user) {
        res.status(400).send({
          message: "The email address is registered."
        });
        return;
      }

      next();
    });
};

const verifySignUp = {
  checkDuplicateEmail: checkDuplicateEmail
};

module.exports = verifySignUp;