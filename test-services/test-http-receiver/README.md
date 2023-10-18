# TEST-HTTP-RECEIVER
test-http-receiver is a service to test http requests. It handles GET, POST, PUT
and DELETE requests and simply logs all information about the request incl. request
body.

## Development
This service is developed using Visual Studio Code and requires the following extensions:
* Docker
* Remote-Containers
* Go

## Deployment
This command runs the service on port 8080.
```
docker run -d -p 8080:8080 --name test-h