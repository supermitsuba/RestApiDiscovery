package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

var locationOfFile = "data/apis.json"

func Get_RestApiRecords(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	var location = r.URL.Query()["location"]
	var environment = r.URL.Query()["environment"]
	var q = r.URL.Query()["q"]
	//var detail = r.URL.Query()["detail"]             // default is return summary, not detail
	var page = r.URL.Query()["page"]                 // default is 0
	var totalRecords = r.URL.Query()["totalRecords"] // default is 100
	var isActive = r.URL.Query()["isActive"]         // default is all, true is all active and false is all inactive

	Get_RestApiRecords_Impl(location, environment, q, page, totalRecords, isActive, w)
}

func Get_RestApiRecords_Impl(location []string,
	environment []string,
	q []string,
	page []string,
	totalRecords []string,
	isActive []string,
	w http.ResponseWriter) {

	if w == nil {
		return
	}

	var restApis = GetFileOfRestApiDescriptions(locationOfFile)

	for j := 0; j < len(location); j++ {
		fmt.Printf("Where Location == %v\n", location[j])
		restApis = Where(restApis, func(item RestApiDescription) bool { return strings.Contains(item.Location, location[j]) })
	}

	for j := 0; j < len(environment); j++ {
		fmt.Printf("Where Environment == %v\n", environment[j])
		restApis = Where(restApis, func(item RestApiDescription) bool { return strings.Contains(item.Environment, environment[j]) })
	}

	for j := 0; j < len(q); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool {
			var toLower = strings.ToLower(q[j])
			fmt.Printf("Where Description(%v) == %v\n", item.Description, q[j])
			fmt.Printf("Where Name(%v) == %v\n", item.Name, q[j])
			return !IsStringWhiteSpaceOrEmpty(q[j]) &&
				(strings.Contains(strings.ToLower(item.Description), toLower) || strings.Contains(strings.ToLower(item.Name), toLower))
		})
	} // default should be everything

	var pageNumber = 0

	if len(page) > 0 {
		var temp, _ = strconv.ParseInt(page[0], 10, 32)
		pageNumber = int(temp)
	}

	var totalRecordsNumber = 100

	if len(totalRecords) > 0 {
		var temp, _ = strconv.ParseInt(totalRecords[0], 10, 32)
		totalRecordsNumber = int(temp)
	}

	restApis = Skip(restApis, pageNumber*totalRecordsNumber)
	restApis = Take(restApis, totalRecordsNumber)

	for j := 0; j < len(isActive); j++ {
		restApis = Where(restApis, func(item RestApiDescription) bool {
			var isBool, err = strconv.ParseBool(isActive[j])
			return err != nil && item.IsActive == isBool
		})
	}

	json.NewEncoder(w).Encode(restApis)
}

func Post_RestApiRecords(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	var x = new(RestApiDescription)
	json.NewDecoder(r.Body).Decode(x)

	Post_RestApiRecords_Impl(x, w)
}

func Post_RestApiRecords_Impl(x *RestApiDescription, w http.ResponseWriter) {
	if w == nil {
		return
	}

	var xList = GetFileOfRestApiDescriptions(locationOfFile)
	x.Id = uuid.New()
	fmt.Printf("Inserting: %v\n", x)
	xList = append(xList, *x)
	fmt.Printf("New List: %v\n", xList)
	WriteRestApiDescriptionsToFile(xList, locationOfFile)
	json.NewEncoder(w).Encode("{ 'result':'OK' }")
}

func Put_RestApiRecords(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	var x = new(RestApiDescription)
	json.NewDecoder(r.Body).Decode(x)
	x.Id = GetId(r)
	Put_RestApiRecords_Impl(x, w)
}

func Put_RestApiRecords_Impl(x *RestApiDescription, w http.ResponseWriter) {
	if w == nil {
		return
	}

	var xList = GetFileOfRestApiDescriptions(locationOfFile)
	var index = Find(xList, x.Id)

	if index == -1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("{ 'result':'Not OK' }")
	} else {
		xList[index] = *x
		WriteRestApiDescriptionsToFile(xList, locationOfFile)
		json.NewEncoder(w).Encode("{ 'result':'OK' }")
	}
}

func Delete_RestApiRecord(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	var id = GetId(r)
	fmt.Printf("Id: %v\n", id)
	Delete_RestApiRecord_Impl(id, w)
}

func Delete_RestApiRecord_Impl(id string, w http.ResponseWriter) {
	if w == nil {
		return
	}
	var xList = GetFileOfRestApiDescriptions(locationOfFile)
	var index = Find(xList, id)

	if index == -1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("{ 'result':'Not OK' }")
	} else {
		xList = Remove(xList, index)
		WriteRestApiDescriptionsToFile(xList, locationOfFile)
		json.NewEncoder(w).Encode("{ 'result':'OK' }")
	}
}

func GetId(r *http.Request) string {
	fmt.Printf("URL: %v\n", r.URL)
	return mux.Vars(r)["id"]
}

func Convert(value []RestApiDescription, item RestApiDescription) []RestApiDescription {
	var slice = make([]RestApiDescription, 0, 5)
	for j := 0; j < len(value)-1; j++ {
		slice = append(slice, value[j])
	}

	slice = append(slice, item)
	return slice
}
