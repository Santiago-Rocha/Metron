package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
	"../service"
)

type OptionHandler struct {
	service.IoptionService
}

func (handler *OptionHandler) CreateOption(res http.ResponseWriter, req *http.Request) {
	var option data.Option

	err := json.NewDecoder(req.Body).Decode(&option)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.InsertOption(option)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "txt :%+v", option)

}
