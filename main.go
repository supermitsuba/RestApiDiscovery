package main

import (
	"RestApiDiscovery/controller"
	"log"
	"net/http"
)

func main() {

	router := controller.NewRouter()

	log.Fatal(http.ListenAndServe(":8088", router))
}
