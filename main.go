package main

import (
	"RestApiDiscovery/libs/controller"
	"log"
	"net/http"
)

func main() {

	router := controller.NewRouter()

	log.Fatal(http.ListenAndServe(":8088", router))
}
