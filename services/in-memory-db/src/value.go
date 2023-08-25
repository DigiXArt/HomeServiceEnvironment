/*
value.go
Implements a single key/value instance, which is saved to the key/value storage.
It also takes care of always setting the right remaining expire time every time
a value is served via the API.

#########################################