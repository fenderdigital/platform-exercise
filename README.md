# Fender Digital Platform Engineering Challenge - EKWONG

## Description

Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Requirements

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
 **Database**:
* Operations:
    * start - npm run startdb
    * stop - npm run stopdb
            
**Node Service**:
* scripts:
    * npm run app - runs express service
    
**Testing**:
* test scripts are saved in *.test.js format in test folder:
    * npm run test - runs defined tests in test folder
    * html report available in test_reports/*.html folder


