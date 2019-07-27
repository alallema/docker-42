Create simple api in go
Redis Database
Datacollector in python

## API Resources

- GET /pulse — heartbeat check if our API is online
- GET /users — fetch all users from the database
- GET /users/[email] — fetch user by email from the database

## Instructions to launch Redis and Api and Python script

In order to bring up:
- Use `docker-compose build` to build api
- Use `docker-compose up` to see the logs of all the containers
- Use `docker-compose up -d` if you want it to run in the foreground
- Launch redis-shell `docker exec -it redis redis-cli`
- Check Api on http://localhost:8080/pulse
- In order to clean up the cluster, use `docker-compose down`

## Dependencies

### Golang
* go-redis
* gorilla

### Python
* pandas
* redis

## Request & Response Examples

### GET /pulse

Example: http://localhost:8080/pulse

Response body:

    {   "status": "OK",
        "code": "200
    }

### GET /users

Example: http://localhost:8080/users

Response body:

    {
      "total":100,
      "users": [
        {
          "lastName":"Hurst",
          "firstName":"Harper",
          "email":"email:Cum.sociis.natoque@eratvolutpatNulla.edu",
          "phoneNumber":"06 22 57 43 11"
        },
        {
          "lastName":"Lara",
          "firstName":"Bradley",
          "email":"email:vulputate.dui.nec@scelerisquenequesed.edu",
          "phoneNumber":"06 90 77 82 50"
        }
      ]
    }

### GET /users/{email}

Example: http://localhost:8080/users/Cum.sociis.natoque@eratvolutpatNulla.edu

Response body:

    {
        "user": {
          "lastName":"Hurst",
          "firstName":"Harper",
          "email":"email:Cum.sociis.natoque@eratvolutpatNulla.edu",
          "phoneNumber":"06 22 57 43 11"
        }
    }

## Redis Structure

HASH    key field value

```
KEY     email:19461183
FIELD   1) "lastName"
        2) "Marshall"
        3) "firstName"
        4) "Mufutau"
        5) "Email"
        6) "porttitor.scelerisque@maurisidsapien.edu"
        7) "phoneNumber"
        8) "09 39 97 87 93"
```
