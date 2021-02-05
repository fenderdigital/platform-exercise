
var chai = require('chai');
var chaiHttp = require('chai-http');
var server = require('../app/index');
var should = chai.should();
const constants = require('../app/constants')
chai.use(chaiHttp);

describe('Route', () => {
  describe('Route Test', () => {
      it('it should GET default route', (done) => {
        chai.request(server)
            .get('/')
            .end((err, res) => {
                  res.should.have.status(200);
                  res.body.should.be.an('object');
                  res.body.should.contain({message:constants.DISPLAY_MESSAGE});
              done();
            });
      });
  });
});