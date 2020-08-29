package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>METRON API IS RUNNING</h1>")
	_ = GetDataBaseClient()
}

func createCategoryHandler(res http.ResponseWriter, req *http.Request) {
	var category category
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Cat :%+v", category)
}

func createTransaction(res http.ResponseWriter, req *http.Request){
	var transaction transaction
	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", transaction)
}
