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
	"os"
	"path/filepath"
	"sync"
	"time"
)

//StorageInterface defines the interface for the data storage.
type StorageInterface interface {
	Initialize(dataRootDirectory string)
	ReadData(collectionName string, startDate time.Time, endDate time.Time) ([]*Data, error)
	WriteData(collectionName string, payload map[string]interface{}) (*Data, error)
	ListCollections() ([]string, error)
}

//Implements StorageInterface
type Storage struct {
	DataRootDirectory string
	MutexLock         sync.Mutex
}

// Initialize sets the data root directory
func (s *Storage) Initialize(dataRootDirectory string) {
	s.DataRootDirectory = dataRootDirectory
}

// getCollectionPath get's the actual path of a collection.
// If the collection does not exist it will create a new directory for the
// colelction first.
func (s *Storage) getCollectionPath(collectionName string) (string, error) {
	path := filepath.Join(s.DataRootDirectory, collectionName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return path, err
		}
	}

	return path, nil
}

// getCurrentDatafilePath gets the path of the current data file
func (s *Storage) getCurrentDatafilePath(collectionName string) (string, error) {
	collectionPath, err := s.getCollectionPath(collectionName)
	if err != nil {
		return "", err
	}

	currentDataFileName := fmt.Sprintf("%v.%v", time.Now().UTC().Format("2006-01-02"), "json")

	return filepath.Join(collectionPath, currentDataFileName), nil
}

// getDataFilePathsInRange get's all path of the data files of a collection
// used in a given time range.
func (s *Storage) getDataFilePathsInRange(collectionName string, startDate time.Time, endDate time.Time) ([]string, error) {
	dataFilePaths := make([]string, 0)
