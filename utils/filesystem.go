package utils

import (
	"os"
)

func MakeDir(path string) error  {
	return os.MkdirAll(path, os.ModePerm)
}