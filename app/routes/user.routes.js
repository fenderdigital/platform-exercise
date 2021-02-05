/*
 *  User End points
*/
const { authToken } = require("../middleware");
const controller = require("../controllers/user.controller");

module.exports = function(app) {
  app.use(function(req, res, next) {
    res.header(
      "Access-Control-Allow-Headers",
      "x-access-token, Origin, Content-Type, Accept"
    );
    next();
  });

//Default route
  app.get("/", controller.allAccess);

//Access protected resource
  app.get(
    "/api/test/user",
    [authToken.verifyToken],
    controller.userPrivate
  );
  app.put(
    "/api/admin/user",
    [authToken.verifyToken],
    controller.nameUpdate
  );
};