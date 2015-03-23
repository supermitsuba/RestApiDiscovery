package main

import (
	"log"
	"net/http"
	"restapidiscovery/controller"
)

func main() {

	router := controller.NewRouter()

	log.Fatal(http.ListenAndServe(":8088", router))
}
