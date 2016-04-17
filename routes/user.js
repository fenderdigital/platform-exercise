var express = require('express');
var middleware = require("./middleware");
var userController = require("../controllers/userController");


var router = express.Router();

router.get("/", middleware.protected, userController.all);
router.get("/:id", middleware.protected, userController.get);
router.put("/:id", middleware.protected, userController.set);
router.delete("/:id", middleware.protected, userController.delete);

module.exports = router;
