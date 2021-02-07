# Fender Digital Platform Engineering Challenge - EKWONG


## Description
Design and implement a RESTful web service to facilitate a user authentication system. 
The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Proposed AWS Architecture
[Proposed AWS Architecture (Simple)](fdr.png)

## Demo Video
[![Platform Exercise Demo Video](http://img.youtube.com/vi/YdtFnYNu7dk/0.jpg)](http://www.youtube.com/watch?v=YdtFnYNu7dk "Platform Exercise Demo Video")

* This project provides a basic token authentication REST API service using node and Express. A working docker container is [built](#quick-start) and [tested](#quick-start).
* Container can be deployed in ECS along with the other AWS components.

## TODO 
* Role based user schema
* move token blacklist from local cache to ElasticCache
* set up Api Gateway, CloudFront and ELB
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

**Quick Start (Ensure System requirements have been met**:

1. Clone this project
2. **echo $PGPASS $PGTEST** : _to verify that password and test flag has been set._
2. change to project directory
3. **npm install && npm run build** : _to install dependencies and build container_
3. **npm run fullTest** : _start database and runs all tests. Report available in test_reports directory_
4. Other scripts :
    * **npm run service-front** : _Start and run all services with console logging_
    * **pm run service** : _Start and run all services in background_
    *** npm run stopdb** : _to stop services._ 

 **Database**:
* Operations:
    * start - npm run startdb
    * stop - npm run stopdb
            
**Node Service**:
* scripts:
    * npm run app - runs express service
    
**Testing**:
* Postman was used for initia tests, included in test/FENDER.postman_collection.json
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

**AWS Deployment**
* Requirements :
    * AWS account
    * [aws cli]("https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html") installed 
    * [** <your account id>** ]("https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html#FindingYourAWSId").dkr.ecr. **<your region>** .amazonaws.com
    * set account id and region in format above
* Log in to ecr : 
    * aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin [** <your account id>** ]("https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html#FindingYourAWSId").dkr.ecr. **<your region>** .amazonaws.com
* Create new ECR repository
    * aws ecr create-repository --repository-name **<repo name>** --region **<your region>** 
        * response output contains key: **repositoryUri**, save that value
        * example:
        * {
    "repository": {
        "repositoryArn": "**arm value**",
        "registryId": "**registry id**",
        "repositoryName": "**repo name**",
        "repositoryUri": "**url**",
       ..
    }
}
* Tag image
    * docker images - get IMAGE ID of your image
    * docker tag **imageid** **repositoryUrl**
* Deploy to ECR
    * docker push **repositoryUrl**
* Create ECS Cluster configuration
    * ecs-cli configure --cluster fender-auth --default-launch-type FARGATE --config-name fender-config --region **<your region>**
    * configure ecs-params.yml with subnet id and security group
    * Create Task
        * ecs-cli compose --project-name fender-auth-service --file docker-compose.yml --ecs-params ecs-params.yml --region **<your region>** create --launch-type FARGATE
    * Create ECS
        * ecs-cli up --cluster-config  fender-auth --vpc **<vpc id>** --subnets **<subnet id** --security-group **<security group id**
    * Start ECS
        * ecs-cli compose --project-name fender-auth-service up --force-deployment --create-log-groups --cluster-config  fender-auth
    
**Troubleshooting**:
* to restart database 
    * npm run stopdb && npm run startdb
* Postgres authentication error
    * ensure environmental variables are set - look at [environment variables](#environment-variable) section above.
    * try running commands with environment variable. example:
        * PGPASS=thepassword PGTEST=true npm run startdb 
        * PGPASS=thepassword PGTEST=true npm run test 
        * PGPASS=thepassword PGTEST=true npm run stopdb 
        




