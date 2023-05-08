
# IN-MEMORY-DB
in-memory-db is a service that does something like redis on a very basic
level.
It implements a very basic key/value storage that can be used to store
data that does not need to be persistet, because the service doesn't do that,
but has to be saved and loaded fast. It is also implemented to automatically
delete data based on an expiration time. There is no way to store data permanently!
Data is lost either after the service restarts or after the set expiration time
is over.
To structure data bit better it implements realms, which is just one layer more
to devide data into seperate spaces. This can be used to seperate storage spaces
for services using this, to eliminate the problem of key conflicts.
You can use this if you want a very lightweight in-memory key/value storage