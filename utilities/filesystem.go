package utilities

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func MakeDir(path string) error  {
	return os.MkdirAll(path, os.ModePerm)
}

func ReadListDir(path string) []fs.FileInfo  {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}