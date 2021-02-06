const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const app = express();
const db = require("./models");
db.sequelize.sync({force: false});
const userRoutes = "./routes/user.routes"
const authRoutes = "./routes/auth.routes"

var corsOptions = {
  origin: "http://localhost:8088"
};

app.use(cors(corsOptions));

app.use(bodyParser.json());

app.use(bodyParser.urlencoded({ extended: true }));

require(authRoutes)(app);
require(userRoutes)(app);

module.exports=app;