
/*
 * Auth End points
*/
const { verifySignUp, authToken } = require("../middleware");
const controller = require("../controllers/auth.controller");

module.exports = function(app) {
  app.use(function(req, res, next) {
    res.header(
      "Access-Control-Allow-Headers",
      "x-access-token, Origin, Content-Type, Accept"
    );
    next();
  });

  app.post(
    "/api/auth/signup",
    [
      verifySignUp.checkDuplicateEmail
    ],
    controller.signup
  );
 app.post("/api/auth/signin", controller.signin);

 //Ensure Front End invalidates token
 app.post("/api/auth/logout",  [authToken.verifyToken,authToken.destroyToken]);
};