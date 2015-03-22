package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

var locationOfFile = "data/apis.json"
var OkResponse = "{ 'result':'OK' }"
var NotFoundResponse = "{ 'result':'Not OK' }"

func Get_RestApiRecords(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var location = request.URL.Query()["location"]
	var environment = request.URL.Query()["environment"]
	var q = request.URL.Query()["q"]
	//var detail = r.URL.Query()["detail"]             // default is return summary, not detail
	var page = request.URL.Query()["page"]                 // default is 0
	var totalRecords = request.URL.Query()["totalRecords"] // default is 100
	var isActive = request.URL.Query()["isActive"]         // default is all, true is all active and false is all inactive

	Get_RestApiRecords_Impl(location, environment, q, page, totalRecords, isActive, response)
}

func Get_RestApiRecords_Impl(location []string,
	environment []string,
	q []string,
	page []string,
	totalRecords []string,
	isActive []string,
	response http.ResponseWriter) {

	if response == nil {
		return
	}

	var restApis = GetFileOfRestApiDescriptions(locationOfFile)
	var totalRecordsNumber = 100
	var pageNumber = 0

	for j := 0; j < len(location); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool { return strings.Contains(item.Location, location[j]) })
	}

	for j := 0; j < len(environment); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool { return strings.Contains(item.Environment, environment[j]) })
	}

	for j := 0; j < len(q); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool {
			var toLower = strings.ToLower(q[j])
			return !IsStringWhiteSpaceOrEmpty(q[j]) &&
				(strings.Contains(strings.ToLower(item.Description), toLower) || strings.Contains(strings.ToLower(item.Name), toLower))
		})
	}

	for j := 0; j < len(isActive); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool {
			var isBool, err = strconv.ParseBool(isActive[j])
			return err != nil && item.IsActive == isBool
		})
	}

	if len(page) > 0 {
		var temp, _ = strconv.ParseInt(page[0], 10, 32)
		pageNumber = int(temp)
	}

	if len(totalRecords) > 0 {
		var temp, _ = strconv.ParseInt(totalRecords[0], 10, 32)
		totalRecordsNumber = int(temp)
	}

	restApis = Skip(restApis, pageNumber*totalRecordsNumber)
	restApis = Take(restApis, totalRecordsNumber)

	json.NewEncoder(response).Encode(restApis)
}

func Post_RestApiRecords(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var item = new(RestApiDescription)
	json.NewDecoder(request.Body).Decode(item)

	Post_RestApiRecords_Impl(item, response)
}

func Post_RestApiRecords_Impl(item *RestApiDescription, response http.ResponseWriter) {
	if response == nil {
		return
	}

	var listOfApis = GetFileOfRestApiDescriptions(locationOfFile)
	item.Id = uuid.New()
	listOfApis = append(listOfApis, *item)
	WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
	json.NewEncoder(response).Encode(OkResponse)
}

func Put_RestApiRecords(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var item = new(RestApiDescription)
	json.NewDecoder(request.Body).Decode(item)
	item.Id = GetId(request)
	Put_RestApiRecords_Impl(item, response)
}

func Put_RestApiRecords_Impl(item *RestApiDescription, response http.ResponseWriter) {
	if response == nil {
		return
	}

	var listOfApis = GetFileOfRestApiDescriptions(locationOfFile)
	var index = Find(listOfApis, item.Id)

	if index == -1 {
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(NotFoundResponse)
	} else {
		listOfApis[index] = *item
		WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
		json.NewEncoder(response).Encode(OkResponse)
	}
}

func Delete_RestApiRecord(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var id = GetId(request)
	Delete_RestApiRecord_Impl(id, response)
}

func Delete_RestApiRecord_Impl(id string, response http.ResponseWriter) {
	if response == nil {
		return
	}
	var listOfApis = GetFileOfRestApiDescriptions(locationOfFile)
	var index = Find(listOfApis, id)

	if index == -1 {
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(NotFoundResponse)
	} else {
		listOfApis = Remove(listOfApis, index)
		WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
		json.NewEncoder(response).Encode(OkResponse)
	}
}

func GetId(r *http.Request) string {
	return mux.Vars(r)["id"]
}
