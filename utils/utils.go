package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Policy struct {
	Site           string   `json:"site"`
	Domains        []string `json:"domains"`
	LinkClasses    []string `json:"link-classes"`
	AllowBacklinks string   `json:"allow-backlinks"`
	Snippets       string   `json:"snippets"`
	SnippetClass   string   `json:"snippet-class"`
	Seed           string   `json:"seed"`
}

// TODO: Refactor the saveToDisk function
func SaveToDisk(name string, body string, dir string) {
	_filename := name + ".html"
	path := filepath.Join("./", dir, _filename)

	os.Mkdir(dir, 0666)
	os.Create(path)
	os.WriteFile(dir+"/"+_filename, []byte(body), 0666)
}

func ReadJson(path string) map[string]Policy {
	/**
	 *Check if the path passed is valid and throw error if not
	 *Parse the json and return it as an object
	**/
	pol := make(map[string]Policy)

	data, err := os.ReadFile(path)
	CheckError(err)

	CheckError(json.Unmarshal(data, &pol))
	return pol
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
