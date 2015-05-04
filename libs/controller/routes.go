package controller

import (
	d "RestApiDiscovery/libs/data"
	interfaces "RestApiDiscovery/libs/data/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var h = Init(interfaces.Data_access(d.File_access{FileLocation: "data.json"}))

func NewRouter() *mux.Router {

	var routes = Routes{
		Route{
			"GetAllRestApiRecords",
			"GET",
			"/api/restapirecords",
			h.Get_RestApiRecords,
		},
		Route{
			"CreateNewRestApiRecord",
			"POST",
			"/api/restapirecords",
			h.Post_RestApiRecords,
		},
		Route{
			"UpdateRestApiRecord",
			"PUT",
			"/api/restapirecords/{id}",
			h.Put_RestApiRecords,
		},
		Route{
			"DeleteRestApiRecord",
			"DELETE",
			"/api/restapirecords/{id}",
			h.Delete_RestApiRecord,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
