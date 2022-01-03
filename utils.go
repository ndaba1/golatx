package main

import (
	"os"
	"path/filepath"
)

func saveToDisk(name string, body string, dir string) {
	_filename := name + ".html"
	path := filepath.Join("./", dir, _filename)

	os.Mkdir(dir, 0666)
	os.Create(path)
	os.WriteFile(dir+"/"+_filename, []byte(body), 0666)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
