var express = require('express');
var userController = require("../controllers/userController");

var router = express.Router();

var port = process.env.PORT || 8080;

router.get('/', function (req, res) {
    res.send('Hello! The API is at http://localhost:' + port + '/api');
});

router.post('/authenticate', userController.authenticate);

module.exports = router;
