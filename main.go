package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./pkg/api"
)

func main() {
	transactionHanlder := ServiceContainer().InjectTransactionController()
	categoryHandler := ServiceContainer().InjectCategoryController()
	router := mux.NewRouter()
	router.HandleFunc("/", api.HomeHandler)
	router.HandleFunc("/categroy/create", categoryHandler .CreateCategory).Methods("POST")
	router.HandleFunc("/transaction/create", transactionHanlder.CreateTransaction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
