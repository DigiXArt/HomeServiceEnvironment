/*
value.go
Implements a single key/value instance, which is saved to the key/value storage.
It also takes care of always setting the right remaining expire time every time
a value is served via the API.

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

import (
	"encoding/json"
	"io"
	"time"
)

// Values implements a single key/value instance, which is saved to the key/value
// storage.
type Value struct {
	Value     string
	ExpiresAt time.Time
}

//ToValueMessageType transforms a Value instance to a ValueMessageType that can
//be converted to json and served via the api. It also takes care of setting the
//right remaining expire time in seconds.
func (v *Value) ToValueMessageType() ValueMessageType {
	return ValueMessageType{
		Value:     v.Value,
		ExpiresIn: (int)(v.ExpiresAt.Sub(time.Now().UTC()).Seconds()),
	}
}

//ValueFromValueMessageType creates a Value from the JSON message in the
//request body and converts the given seconds into a time instance.
func ValueFromValueMessageType(body io.ReadCloser) (*Value, error) {
	msg := ValueMessageType{}

	err := json.NewDecoder(body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	value := &Value{
		Value:     msg.Value,
		ExpiresAt: time.Now().UTC().Add(time.Duration(msg.ExpiresIn) * time.Second),
	}

	return value, nil
}
