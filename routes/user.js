var express = require('express');
var middleware = require("./middleware");
var userController = require("../controllers/userController");


var router = express.Router();

router.get('/', middleware.protected, userController.all);

module.exports = router;
