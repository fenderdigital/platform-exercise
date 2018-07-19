# Fender Digital Platform Engineering Challenge

## Description

Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.
Given these requirements, JWT makes sense.  The implementation here could be improved by incorporating middleware token handling.

## Requirements

**Models**

1. user (implemented)
2. session (implemented)
3. password_history (not implemented)
4. user_record_history (not implemented)

**Endpoints**

All of these endpoints should be written from a user's perspective. The token the user receives upon login must be provided as an Authorization header for Logout, Update, and Delete operations.

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

Please follow these steps to configure:
- Update the jwt signing key in the line 6 of platform-test/src/secret/secret.go, in the method Fetch
- `docker build -t platform_test .` (build the docker image)
- `docker-compose up` (instantiate the containers)
- `docker exec -it postgres_docker_container_id psql -U postgres`
  - Copy the contents of this dbsetup and paste into the postgres prompt

**Test**
Docker exec into the running server container (use `docker ps` to get the container id)
- `docker exec -it server_container_hash sh`
- `go test ./... -v` Run the tests

**Enhancements**
- Fix the plain text password transmission
- Enforce email uniqueness across Users
- Refactor and DRY up things
- Standardize Info and Error logging
- Track password and user record updates
- Add an admin claim that allows reading and modifying records that other than the user's own
- Plan and build integrations for 3rd party login providers like google, facebook, linkedin, github etc
- Compaction function fired from lambda every day that removes sessions records > week old
