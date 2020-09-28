package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./pkg/api"
)

func main() {
	transactionHanlder := ServiceContainer().InjectTransactionController()
	optionHandler := ServiceContainer().InjectOptionController()
	router := mux.NewRouter()
	router.HandleFunc("/", api.HomeHandler)
	router.HandleFunc("/transaction/create", transactionHanlder.CreateTransaction).Methods("POST")
	router.HandleFunc("/option/create", optionHandler.CreateOption).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
