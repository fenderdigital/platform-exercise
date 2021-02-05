const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const app = express();

const userRoutes = "./routes/user.routes"

var corsOptions = {
  origin: "http://localhost:8088"
};

app.use(cors(corsOptions));

app.use(bodyParser.json());

app.use(bodyParser.urlencoded({ extended: true }));

require(userRoutes)(app);

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`);
});

module.exports=app;