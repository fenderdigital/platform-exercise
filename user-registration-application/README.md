# Demo project

CRUD service used for member creation, update, login, log out and deletion

## Setup/Configuration

In order to run this application a postgresDb will need to install and running.

refer to for installation: https://wiki.postgresql.org/wiki/Detailed_installation_guides

once install run postgres server
for mac:

 1. pg_ctl -D /usr/local/var/postgres start
 
 2. psql postgres
 
 3. CREATE DATABASE testdb;




## Installation

From project root:

```bash
$ mvn clean install 
```


## Running the microService



```bash
$ mvn spring-boot:run -Dspring.profile.active=default

```
or 

```bash
$ import in IDE and run application
```


## My approach

I wanted to try to keep things simple, I used springboot to create my application
and I integrated with hibernate for simple database interaction. 
I wanted to use spring security to manage authentication, but ran out of time.
To manage access I used jwt tokens created from a utility class and stored in a in memory cache.
A better approach would have been to store in a database or redis server. I also used the jwt to remove path
parameters from the delete and update methods allowing on the user signed in to delete and update itself.
In the future or if I more time I would add an admin controller that can delete and update accounts passing in an user id.
Also I would add a get user endpoint.
