
const constants = require('../constants')

  exports.allAccess = (req, res) => {
    res.status(200).send({"message": constants.DISPLAY_MESSAGE});
  };
    
  