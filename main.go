package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/create", createCategoryHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
