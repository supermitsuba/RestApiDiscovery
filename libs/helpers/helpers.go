package helpers

import (
	"RestApiDiscovery/libs/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func IsStringWhiteSpaceOrEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func GetId(r *http.Request) string {
	return mux.Vars(r)["id"]
}

func ConvertToThing(value string) []model.RestApiDescription {
	var results = new([]model.RestApiDescription)
	json.Unmarshal([]byte(value), results)
	return *results
}
