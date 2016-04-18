# Fender Digital Platform Engineering Challenge

## Description

Design and implement a RESTful web service to facilitate a user authentication system. The authentication mechanism should be *token based*. Requests and responses should be in **JSON**.

## Usage
### Token

The token is returned by the authenticate Endpoint, the token can be pass using
**x-access-token** header
or as a parameter in the url i.e.
/users?**token=(valid token)**

### Endpoints

GET '/users' **protected**

POST '/users'  **protected**

**Parameters**

email (string) **unique** 
name (string)
password (string)

GET '/users/:id' **protected**

PUT '/users/:id' **protected**

**Parameters**

email (string)
name (string)
password (string)

DELETE '/users/:id' **protected**

POST '/authenticate' 

**Parameters**

email (string)
password (string)

POST '/logout' **protected**

