package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"restapidiscovery/model"
)

// todo: make sure file stream is disposed
func WriteRestApiDescriptionsToFile(list []model.RestApiDescription, fileName string) {
	var listAsString, _ = json.Marshal(list)

	ioutil.WriteFile(fileName, []byte(listAsString), 0777)
}

// todo: dispose of file stream
func GetFileOfRestApiDescriptions(fileName string) []model.RestApiDescription {
	var results = new([]model.RestApiDescription)
	var fileStream, _ = os.Open(fileName)

	json.NewDecoder(fileStream).Decode(results)
	fileStream.Close()
	return *results
}
