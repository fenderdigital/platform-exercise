var config = {
    "development": {
        "dialect": "postgres",
        "url": "postgres://zufzdmqcjtvjno:b0iaIich2Ez1wbp7sGeplvFqFh@ec2-54-204-40-209.compute-1.amazonaws.com:5432/d3d971cqlhefct",
        "secret": 'somethingsomethingdarkside',
        "expiration": 600,
        "saltRounds": 10,
        "redis": "redis://h:p52h46aacr82889dp8h7eg1uoc@ec2-54-227-252-69.compute-1.amazonaws.com:17059"
    },
    "test": {
        "username": "root",
        "password": null,
        "database": "database_test",
        "host": "127.0.0.1",
        "dialect": "mysql"
    },
    "production": {
        "username": "root",
        "password": null,
        "database": "database_production",
        "host": "127.0.0.1",
        "dialect": "mysql"
    }
};

module.exports = config;
