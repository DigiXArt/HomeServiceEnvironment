/*
AUTH

Auth is a service that can be used to authenticate user and retrieve
permissions. Login returns an access and a refresh token. The access token
can be used to autheticate this user in other services, which can call decode
to verify the token. Decode retirns the user's id as well as all permissio