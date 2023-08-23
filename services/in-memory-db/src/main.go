
/*
IN-MEMORY-DB

in-memory-db is a service that does something like redis on a very basic
level.
It implements a very basic key/value storage that can be used to store
data that does not need to be persistet, because the service doesn't od that,
but has to be saved and loaded fast. It is also implemented to automatically
delete data based on an expiration time. There is no way to store data permanently!
Data is lost either after the service restarts or after the set expiration time
is over.
To structure data bit better it implements realms, which is just one layer more
to devide data into seperate spaces. This can be used to seperate storage spaces
for services using this, to eliminate the problem of of key conflicts.
You can use this if you want a very lightweight in-memory key/value storage
and redis is just too much, or use it to see how key/value databases could be
implemented in a very basic way. It also shows how you can use go routines to do
things after a set amount of time asynchronously.

###################################################################################

main.go
This is the main entrypoint of the service. It starts the service and
routes all API methods.

###################################################################################

MIT License

Copyright (c) 2020 Bruno Hautzenberger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is