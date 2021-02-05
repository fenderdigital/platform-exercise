/*
 *
 * Use Node Cache in this example to store blacklisted tokens"
 * TODO: migrate to Regis and delete expired tokens
 *  
 */

const jsonwebtoken = require("jsonwebtoken");
const config = require("../configurations/auth.config.js");
const db = require("../models");
const controller = require("../controllers/auth.controller");
const NodeCache = require( "node-cache" );
const blacklist = new NodeCache();

verifyToken = (req, res, next) => {
  let token = req.headers["x-access-token"];

  if ((!token) || (token&&blacklist.get(token))) {
    return res.status(403).send({
      message: "Unauthorized Access Forbidden!"
    });
  }
  jsonwebtoken.verify(token, config.secret, (err, decoded) => {
    if (err) {
      return res.status(401).send({
        message: "Unauthorized Access Forbidden!"
      });
    }
    req.userId = decoded.id;
    controller.validateUser(req, res, next)
 
  });
};


 
destroyToken= (req, res, next) => {
  let token = req.headers["x-access-token"];

  blacklist.set( token, token )

  return res.status(200).send({
    message: "Logged out"
  });
};


const authToken = {
  verifyToken: verifyToken,
  destroyToken: destroyToken
};
module.exports = authToken;