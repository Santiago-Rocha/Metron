package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
	"../service"
)

type TransactionHandler struct {
	service.ItransactionService
}

func (handler *TransactionHandler) CreateTransaction(res http.ResponseWriter, req *http.Request) {
	var transactionDTO data.TransactionDTO
	err := json.NewDecoder(req.Body).Decode(&transactionDTO)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.InsertTransaction(transactionDTO)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", transactionDTO)
}
