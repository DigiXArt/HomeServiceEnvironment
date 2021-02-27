/*
AUTH

Auth is a service that can be used to authenticate user and retrieve
permissions. Login returns an access and a refresh token. The access token
can be used to autheticate this user in other services, which can call decode
to verify the token. Decode retirns the user's id as well as all permissions
of this user. The refresh token can be used to refresh a user'S authentication.
Again, do not use this in production, but it is a nice example on how to
implment JWT authentication and a refresh mechanism.

###################################################################################

main.go
This is the main entrypoint of the service. It starts the service and
routes all API methods.

###################################################################################

MIT License

Copyright (c) 2020 Bruno Hautzenberger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without lim