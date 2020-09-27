package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
	"../service"
)

type CategoryHandler struct {
	service.IcategoryService
}

func (handler *CategoryHandler) CreateCategory(res http.ResponseWriter, req *http.Request) {
	var category data.Category
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.InsertCategory(category)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", category)
}
