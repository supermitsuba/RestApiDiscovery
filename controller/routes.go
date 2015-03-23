package controller

import (
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

func NewRouter() *mux.Router {

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

var routes = Routes{
	Route{
		"GetAllRestApiRecords",
		"GET",
		"/api/restapirecords",
		Get_RestApiRecords,
	},
	Route{
		"CreateNewRestApiRecord",
		"POST",
		"/api/restapirecords",
		Post_RestApiRecords,
	},
	Route{
		"UpdateRestApiRecord",
		"PUT",
		"/api/restapirecords/{id}",
		Put_RestApiRecords,
	},
	Route{
		"DeleteRestApiRecord",
		"DELETE",
		"/api/restapirecords/{id}",
		Delete_RestApiRecord,
	},
}
