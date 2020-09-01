package service

import (
	"../data"
	"../repository"
)

type transactionService struct {
	repository.ItransactionRepository
}

func (service *transactionService) InsertTransaction(transaction data.Transaction) {
	service.InsertTransaction(transaction)
}

func (service *transactionService) GetTransaction(id string) (data.Transaction, error) {
	return service.GetTransaction(id)
}
