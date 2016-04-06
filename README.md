# Fender Digital Platform Engineering Challenge

## Description

Fork this repo.  Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Requirements


**Models**

The **User** model should have the following properties (at minimum):

`name`

`email`

`password`

You should determine what, *if any*, additional models you will need.

**Endpoints**

All of these endpoints should be written from a user's perspective.

1. **User** Registration
2. Login (*token based*) - should return a token, given *valid* credentials
3. Logout - logs a user out
4. Update a **User**'s Information
5. Delete a **User**

**Additional Info**

We expect this project to take an hour or two to complete (based on Rails or Sinatra, but other options include Python or Go). Feel free to use whichever database you'd like; we suggest Postgres. Bonus points for security, specs, etc. Do as little or as much as you like.

Please commit your code into your forked git repo, and show your work and process through those commits.

