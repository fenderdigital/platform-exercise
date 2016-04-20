var config = {
    "development": {
        "dialect": "postgres",
        "url": "postgres://agmfopzvgqmqcr:hcZM_giWimnCiYodPPmucyu9qT@ec2-174-129-37-54.compute-1.amazonaws.com:5432/da98j3iuovg4cr",
        "secret": 'somethingsomethingdarkside',
        "expiration": 600,
        "saltRounds": 10,
        "redis": "redis://h:prjd1ak1553m7deh178pph3gsf@ec2-54-227-252-91.compute-1.amazonaws.com:13609"
    },
    "test": {
        "dialect": "postgres",
        "url": "postgres://zufzdmqcjtvjno:b0iaIich2Ez1wbp7sGeplvFqFh@ec2-54-204-40-209.compute-1.amazonaws.com:5432/d3d971cqlhefct",
        "secret": 'somethingsomethingdarkside',
        "expiration": 600,
        "saltRounds": 10,
        "redis": "redis://h:p52h46aacr82889dp8h7eg1uoc@ec2-54-227-252-69.compute-1.amazonaws.com:17059"
    },
    "production": {
        "dialect": "postgres",
        "url": "postgres://zufzdmqcjtvjno:b0iaIich2Ez1wbp7sGeplvFqFh@ec2-54-204-40-209.compute-1.amazonaws.com:5432/d3d971cqlhefct",
        "secret": 'somethingsomethingdarkside',
        "expiration": 600,
        "saltRounds": 10,
        "redis": "redis://h:p52h46aacr82889dp8h7eg1uoc@ec2-54-227-252-69.compute-1.amazonaws.com:17059"
    }
};

module.exports = config;
