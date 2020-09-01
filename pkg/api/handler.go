package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
	"../db"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>METRON API IS RUNNING</h1>")
	_ = db.NewDatastore()
}

func CreateCategory(res http.ResponseWriter, req *http.Request) {
	var category data.Category
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Cat :%+v", category)
}

func CreateTransaction(res http.ResponseWriter, req *http.Request) {
	var transaction data.Transaction
	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", transaction)
}

func GetTransactionsByCategory(res http.ResponseWriter, req *http.Request) {
	var transaction data.Transaction
	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", transaction)
}
