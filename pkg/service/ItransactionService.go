package service

import (
	"../data"
)

type ItransactionService interface {
	InsertTransaction(data.TransactionDTO) error
	FindTransaction(id string) (data.Transaction, error)
}
