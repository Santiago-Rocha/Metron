package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./pkg/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HomeHandler)
	router.HandleFunc("/create", api.CreateCategory).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
