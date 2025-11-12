package main

import (
	"encoding/json"
	"io/ioutil"
)

// LoadJSON reads a JSON file into v.
func LoadJSON(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// SaveJSON writes v to path with indentation.
func SaveJSON(path string, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}
