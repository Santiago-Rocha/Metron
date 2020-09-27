package service

import (
	"../data"
	"../repository"
)

type TransactionService struct {
	repository.ItransactionRepository
}

func (service *TransactionService) InsertTransaction(transaction data.Transaction) error {
	return service.Insert(transaction)
}

func (service *TransactionService) GetTransaction(id string) (data.Transaction, error) {
	return service.Find(id)
}
