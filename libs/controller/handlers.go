package controller

import (
	interfaces "RestApiDiscovery/libs/data/interfaces"
	"RestApiDiscovery/libs/helpers"
	"RestApiDiscovery/libs/model"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handlers struct {
	OkResponse       string
	NotFoundResponse string
	DataAccess       interfaces.Data_access
}

func Init(data_layer interfaces.Data_access) *Handlers {
	dataLayer := new(Handlers)
	dataLayer.DataAccess = data_layer
	dataLayer.OkResponse = "{ 'result':'OK' }"
	dataLayer.NotFoundResponse = "{ 'result':'Not OK' }"
	return dataLayer
}

func (h Handlers) Get_RestApiRecords(response http.ResponseWriter, request *http.Request) {
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

	h.Get_RestApiRecords_Impl(location, environment, q, page, totalRecords, isActive, response)
}

func (h Handlers) Get_RestApiRecords_Impl(location []string,
	environment []string,
	q []string,
	page []string,
	totalRecords []string,
	isActive []string,
	response http.ResponseWriter) {

	if response == nil {
		return
	}

	var restApis = helpers.ConvertToThing(h.DataAccess.Load("")) //  data.GetFileOfRestApiDescriptions(h.locationOfFile)
	var totalRecordsNumber = 100
	var pageNumber = 0

	for j := 0; j < len(location); j++ {
		restApis = model.Where(restApis, func(item model.RestApiDescription) bool { return strings.Contains(item.Location, location[j]) })
	}

	for j := 0; j < len(environment); j++ {
		restApis = model.Where(restApis, func(item model.RestApiDescription) bool { return strings.Contains(item.Environment, environment[j]) })
	}

	for j := 0; j < len(q); j++ {
		restApis = model.Where(restApis, func(item model.RestApiDescription) bool {
			var toLower = strings.ToLower(q[j])
			return !helpers.IsStringWhiteSpaceOrEmpty(q[j]) &&
				(strings.Contains(strings.ToLower(item.Description), toLower) || strings.Contains(strings.ToLower(item.Name), toLower))
		})
	}

	for j := 0; j < len(isActive); j++ {
		restApis = model.Where(restApis, func(item model.RestApiDescription) bool {
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

	restApis = model.Skip(restApis, pageNumber*totalRecordsNumber)
	restApis = model.Take(restApis, totalRecordsNumber)

	json.NewEncoder(response).Encode(restApis)
}

func (h Handlers) Post_RestApiRecords(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var item = new(model.RestApiDescription)
	json.NewDecoder(request.Body).Decode(item)

	h.Post_RestApiRecords_Impl(item, response)
}

func (h Handlers) Post_RestApiRecords_Impl(item *model.RestApiDescription, response http.ResponseWriter) {
	if response == nil {
		return
	}

	var listOfApis = helpers.ConvertToThing(h.DataAccess.Load("")) //data.GetFileOfRestApiDescriptions(locationOfFile)
	item.Id = uuid.New()
	listOfApis = append(listOfApis, *item)
	var data, _ = json.Marshal(listOfApis)
	h.DataAccess.Save("", string(data))
	//data.WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
	json.NewEncoder(response).Encode(h.OkResponse)
}

func (h Handlers) Put_RestApiRecords(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var item = new(model.RestApiDescription)
	json.NewDecoder(request.Body).Decode(item)
	item.Id = helpers.GetId(request)
	h.Put_RestApiRecords_Impl(item, response)
}

func (h Handlers) Put_RestApiRecords_Impl(item *model.RestApiDescription, response http.ResponseWriter) {
	if response == nil {
		return
	}

	var listOfApis = helpers.ConvertToThing(h.DataAccess.Load("")) //data.GetFileOfRestApiDescriptions(locationOfFile)
	var index = model.Find(listOfApis, item.Id)

	if index == -1 {
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(h.NotFoundResponse)
	} else {
		listOfApis[index] = *item
		var data, _ = json.Marshal(listOfApis)
		h.DataAccess.Save("", string(data)) //data.WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
		json.NewEncoder(response).Encode(h.OkResponse)
	}
}

func (h Handlers) Delete_RestApiRecord(response http.ResponseWriter, request *http.Request) {
	if response == nil || request == nil {
		return
	}

	var id = helpers.GetId(request)
	h.Delete_RestApiRecord_Impl(id, response)
}

func (h Handlers) Delete_RestApiRecord_Impl(id string, response http.ResponseWriter) {
	if response == nil {
		return
	}
	var listOfApis = helpers.ConvertToThing(h.DataAccess.Load("")) //data.GetFileOfRestApiDescriptions(locationOfFile)
	var index = model.Find(listOfApis, id)

	if index == -1 {
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(h.NotFoundResponse)
	} else {
		listOfApis = model.Remove(listOfApis, index)
		var data, _ = json.Marshal(listOfApis)
		h.DataAccess.Save("", string(data))
		//data.WriteRestApiDescriptionsToFile(listOfApis, locationOfFile)
		json.NewEncoder(response).Encode(h.OkResponse)
	}
}
