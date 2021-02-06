module.exports = {
    HOST: "postgres",
    USER: "postgres",
    PASSWORD: "Tele2021!",
    DB: "testdb",
    dialect: "postgres",
    pool: {
      max: 5,
      min: 0,
      acquire: 30000,
      idle: 10000
    }
  };