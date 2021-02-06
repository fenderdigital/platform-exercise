# Fender Digital Platform Engineering Challenge - EKWONG

## Description

Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Requirements
**Install Node and NPM**:
Windows: https://cloudlinuxtech.com/how-to-install-node-js-npm/
Mac: https://blog.teamtreehouse.com/install-node-js-npm-mac
Linux: https://linuxconfig.org/how-to-install-node-js-on-linux

**Install Docker**: 
Windows: https://docs.docker.com/docker-for-windows/install/
Mac: https://docs.docker.com/docker-for-mac/install/
Linux: https://docs.docker.com/engine/install/ , click on Installations per distro

**Environment Variables**:
* Required environment variables :
    *  PGPASS - password used for database
    * PGTEST - boolean, set to true when running test.

* Setting environment variables:
    * Windows: http://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-windows-command-line-and-registry/
    * Mac: https://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-mac-os-x-terminal-only/
    * Linux: https://linuxize.com/post/how-to-set-and-list-environment-variables-in-linux/
* example : export PGPASS=password 
* verify environment variables:
    * echo $PGPASS : displays value. 

## Running and testing the project

**Installation**:
* clone this project
* run npm install in project directory

**Quick Start**:
* npm run fullTest : to build database and run all tests.
* npm run service-front : Start and run all services with console logging
    * interactive kill on console to end. 
* npm run service : Start and run all services in background
    * npm run stopdb : to stop services. 

 **Database**:
* Operations:
    * start - npm run startdb
    * stop - npm run stopdb
            
**Node Service**:
* scripts:
    * npm run app - runs express service
    
**Testing**:
* test scripts are saved in *.test.js format in test folder:
    * npm run test - runs all tests in test folder
        * signup.test.js - test user sign up
        * signin.test.js - test user sign in
        * logout.test.js - test user sign out
        * route.test.js - test default route
        * auth_access.test.js - test protected resource access
        * auth_update.test.js - test protected resource access
    * npm run fullTest - builds the database and runs all tests, shuts down database
    * html report available in test_reports/*.html folder

**Service**
* To start service, npm run service

**Troubleshooting**:
* to restart database 
    * npm run stopdb && npm run startdb
* Postgres authentication error
    * ensure environmental variables are set - look at [environment variables](#environment-variable) section above.
    * try running commands with environment variable. example:
        * PGPASS=thepassword PGTEST=true npm run startdb 
        * PGPASS=thepassword PGTEST=true npm run test 
        * PGPASS=thepassword PGTEST=true npm run stopdb 
        




