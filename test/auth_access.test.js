
const db = require('../app/models');
db.sequelize.sync({force: false})
const User = db.user;
var assert = require('assert');
var chai = require('chai');
var chaiHttp = require('chai-http');
var server = require('../app/index');
var should = chai.should();
const constants = require('../app/constants')

var token=""

chai.use(chaiHttp);

describe('Protected Resource Access', () => {
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
                  res.body.should.contain({message:constants.REGISTERED_SUCCESS});
              done();
            });
      });
      it('it should fail GET user content without token', (done) => {
        chai.request(server)
            .get('/api/test/user')
            .set('x-access-token', token)
            .end((err, res) => {
                  res.should.have.status(403);
                  res.body.should.be.an('object');
                  res.body.should.contain({message:constants.UNAUTH_FORBIDDEN});
                  done();
            });
      });
      it('it should POST to signin user', (done) => {
        chai.request(server)
            .post('/api/auth/signin')
            .send({email: constants.TEST_USER.email, password: constants.TEST_USER.password})
            .end((err, res) => {
                  res.should.have.status(200);
                  res.body.accessToken.should.be.a('string');
                  token=res.body.accessToken;
                  done();
            });
      });
      it('it should GET user content with token', (done) => {
        chai.request(server)
            .get('/api/test/user')
            .set('x-access-token', token)
            .end((err, res) => {
                  res.should.have.status(200);
                  done();
            });
      });
  });
});