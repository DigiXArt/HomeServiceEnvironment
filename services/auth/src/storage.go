
/*
storage.go
Defines the storage interface of this application and implements the
actual json object storage based on data files.

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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var usersDirectory = "users"
var servicesDirectory = "services"

//StorageInterface defines the interface for the data storage.
type StorageInterface interface {
	Initialize(dataRootDirectory string)
	GetUserByCredentials(username string, passowrd string) (*User, bool, error)
	GetUser(ID string) (*User, error)
	GetServiceByCredentials(ID string, key string) (*Service, bool, error)
	GetService(ID string) (*Service, error)
}

// Storage implements StorageInterface
type Storage struct {
	DataRootDirectory string
}

// Initialize sets the data root directory
func (s *Storage) Initialize(dataRootDirectory string) {
	s.DataRootDirectory = dataRootDirectory
}

// fileExists checks if a file exists
func (s *Storage) fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// getDirectoryPath is used to get a directory inside the data root
// directory. If it does not exist, it will be created.
func (s *Storage) getDirectoryPath(name string) (string, error) {
	path := filepath.Join(s.DataRootDirectory, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return path, err
		}
	}

	return path, nil
}