
const db = require('../app/models');
db.sequelize.sync({force: false})
const User = db.user;
var assert = require('assert');
var chai = require('chai');
var chaiHttp = require('chai-http');
var server = require('../app/index');
var should = chai.should();
const constants = require('./constants')


chai.use(chaiHttp);

describe('User', () => {
  after((done) => {
    User.findOne({
        where: {
          email: constants.TEST_USER.email
        }
      }).then((user)=>{
        if (user){
          user.destroy().then((res)=>{
            done();
          });
        }
      })
  });
  describe('Sign up', () => {
      it('it should POST to signup new user', (done) => {
        chai.request(server)
            .post('/api/auth/signup')
            .send(constants.TEST_USER)
            .end((err, res) => {
                  res.should.have.status(200);
                  res.body.should.be.an('object');
                  res.body.should.contain({message:"User was registered successfully!"});
              done();
            });
      });
  });

});