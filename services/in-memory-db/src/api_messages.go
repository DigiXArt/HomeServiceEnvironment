
/*
api_messages.go
Defines all types that are either served via the API or received from the client.

###################################################################################

MIT License

Copyright (c) 2020 Bruno Hautzenberger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

//ValueMessageType defines the API message for Values
type ValueMessageType struct {
	Value     string `json:"value"`
	ExpiresIn int    `json:"expires-in"`
}

//KeyListMessageType defines the API message for lists of keys
type KeyListMessageType struct {
	Keys []string `json:"keys"`
}

//RealmListMessageType defines the API message for lists of realms
type RealmListMessageType struct {
	Realms []string `json:"realms"`
}

//ErrorMessageType defines the API message for errors
type ErrorMessageType struct {
	Error interface{} `json:"error"`
}