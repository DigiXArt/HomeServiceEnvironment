
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
and redis is just too much, or use it to see how key/value databases could be
implemented in a very basic way. It also shows how you can use go routines to do
things after a set amount of time asynchronously.

This service is be able to:
* Store Data (SET), which will automatically expire.
* Load Data (GET)
* Explicitly delete Data (DELETE)
* List all keys in a realm (LIST-KEYS)
* List all realms (LIST-REALMS)

## Development
This service is developed using Visual Studio Code and requires the following extensions:
* Docker
* Remote-Containers
* Go

## API
Description and examples (cUrl) of all API calls and models of this service.

### Models
#### Value
```json
{
        "value":"a value as string",
        "expires-in":180
}
```

#### Key List
```json
{
        "keys":["key1", "key2", ...]
}