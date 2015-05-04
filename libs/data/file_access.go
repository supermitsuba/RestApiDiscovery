package data

import (
	"io/ioutil"
)

type File_access struct {
	FileLocation string
}

func (r File_access) Save(key string, value string) {
	ioutil.WriteFile(r.FileLocation, []byte(value), 0777)
}

func (r File_access) Load(key string) string {
	var value, _ = ioutil.ReadFile(r.FileLocation)

	return string(value)
}
