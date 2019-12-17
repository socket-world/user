package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Create a router and setup route handlers.
	router := mux.NewRouter().StrictSlash(true)
	router.Methods(`GET`).Path(`/{name}`).HandlerFunc(Get)
	router.Methods(`POST`).Path(`/{name}`).HandlerFunc(Post)

	// Start HTTP server, log error on exit.
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}
