package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// todo: make sure file stream is disposed
func WriteRestApiDescriptionsToFile(list []RestApiDescription, fileName string) {
	var str, _ = json.Marshal(list)
	ioutil.WriteFile(fileName, []byte(str), 0777)
}

// todo: dispose of file stream
func GetFileOfRestApiDescriptions(fileName string) []RestApiDescription {
	var newList = new([]RestApiDescription)
	var file, _ = os.Open(fileName)
	json.NewDecoder(file).Decode(newList)
	file.Close()
	return *newList
}
