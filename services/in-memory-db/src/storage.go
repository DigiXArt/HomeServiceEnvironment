
/*
storage.go
Defines the storage interface interface of this application and implements the
actual in-memory storage based on built-in go maps.

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
	"log"
	"time"
)

//StorageInterface defines the interface for the in-memory key/value storage.
type StorageInterface interface {
	Initialize()
	Get(realmName string, key string) (bool, *Value)
	Set(realmName string, key string, value *Value)
	Delete(realmName string, key string) bool
	Keys(realmName string) []string
	Realms() []string
}

//Implements StorageInterface
type Storage struct {
	Data map[string]map[string]*Value
}

//Initialize creates an empty map[string]map[string]*Value (REALM->KEY->VALUE),
//which will be used to save to store all the data.
func (s *Storage) Initialize() {
	s.Data = make(map[string]map[string]*Value)
}

//GetRealm returns all data of an existing realm as map[string]*Value
func (s *Storage) GetRealm(realm string) (bool, map[string]*Value) {
	if _, ok := s.Data[realm]; ok {
		return true, s.Data[realm]
	}

	return false, make(map[string]*Value)
}

//CreateRealm creates a new realm, if it does not exist already and
//returns it.
func (s *Storage) CreateRealm(realm string) map[string]*Value {
	if _, ok := s.Data[realm]; ok {
		return s.Data[realm]
	}

	s.Data[realm] = make(map[string]*Value)

	return s.Data[realm]
}

//CleanEmptyRealm removes all empty realms, because there is no need to
//keep empty storage spaces.
func (s *Storage) CleanEmptyRealm(realmName string) {
	if realm, ok := s.Data[realmName]; ok {
		if len(realm) == 0 {
			delete(s.Data, realmName)
		}
	}
}

//Get loads a single Value identified by realm and key.
//First bool return value determines, if a Value with these identifiers
//was found, if false Valiue will be nil.
func (s *Storage) Get(realmName string, key string) (bool, *Value) {