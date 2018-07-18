# Fender Digital Platform Engineering Challenge

## Description

Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Requirements

**Models**

1. user (implemented)
2. session (implemented)
3. password_history (not implemented)
4. user_record_history (not implemented)

**Endpoints**

All of these endpoints should be written from a user's perspective.

1. **User** Registration
2. Login (*token based*) - should return a token, given *valid* credentials
3. Logout - logs a user out
4. Update a **User**'s Information
5. Delete a **User**

**README**

Please include:
- a readme file that explains your thinking
- how to setup and run the project
- if you chose to use a database, include instructions on how to set that up
- if you have tests, include instructions on how to run them
- a description of what enhancements you might make if you had more time.

**Setup**

- I'm using docker compose to run this project, however since I choose to load the compiled go file I ran into bit of a rabbit hole setting up the script that waits for the db to start.  In production, I'd be reasonably sure the database was reachable because I would have a separate resilient database cluster.  To mitigate this local, please bring up the database first and then run the web server by executing these commands in order:
- `docker build -t platform_test .`  (build the docker image)
- `docker-compose run -p 5432:5432 -d localhost` (start the postgres server)
- ``

Please fork this repo and commit your code into that fork.  Show your work and process through those commits.
