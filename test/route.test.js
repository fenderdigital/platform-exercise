
var chai = require('chai');
var chaiHttp = require('chai-http');
var server = require('../app/index');
var should = chai.should();

chai.use(chaiHttp);

describe('Route', () => {
  describe('Route Test', () => {
      it('it GET default route', (done) => {
        chai.request(server)
            .get('/')
            .end((err, res) => {
                  res.should.have.status(200);
                  res.body.should.be.an('object');
                  res.body.should.contain({message:"Displaying Content."});
              done();
            });
      });
  });
});