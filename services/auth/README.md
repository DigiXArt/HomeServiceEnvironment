
# AUTH
Auth is a service that can be used to authenticate user and retrieve
permissions. Login returns an access and a refresh token. The access token
can be used to autheticate this user in other services, which can call decode 
to verify the token. Decode retirns the user's id as well as all permissions
of this user. The refresh token can be used to refresh a user'S authentication.
Again, do not use this in production, but it is a nice example on how to
implment JWT authentication and a refresh mechanism.

## Development
This service is developed using Visual Studio Code and requires the following extensions:
* Docker
* Remote-Containers
* Go

## Deployment
This command runs the service on port 7004 and mounts the local directory /media/external/storage/auth to /data which will be used by the service to read user data from. IT also sets all secrets and also the password for the super user (su).
```
docker run -d -p 7004:7004 --name auth -e PORT='7004' -e DATA_DIRECTORY='/data' -e AUTH_ACCESS_SECRET='myGreatAccessSecret' -e AUTH_REFRESH_SECRET='myGreatRefreshSecret' -e AUTH_SERVICE_SECRET='myGreatServiceSecret' -e SU_PWD='myGreatSuPassword' -v /var/run/docker.sock:/var/run/docker.sock --restart unless-stopped --mount type=bind,source=/media/external/storage/auth,target=/data data-logger:1.0
```

## API
Description and examples (cUrl) of all API calls and models of this service

### Errors
All errors are served as an object like this and will return a suitable HTTP
status code.
```json
{
        "error":{
                "message":"Invalid Token",
                "status":403,
                "code":3
                }
}
```

### Methods
#### LOGIN
Logs in a user using username and password.
```
curl --header "Content-Type: application/json" \
        --request POST \ 
        --data '{"username":"theUsername","password":"thePAssword"}' \
        http://localhost:7004/login
```

This call returns a access-token and a refresh token for this session.