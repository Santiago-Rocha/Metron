package service

import (
	"../data"
)

type ItransactionService interface {
	InsertTransaction(data.Transaction) error
	GetTransaction(id string) (data.Transaction, error)
}
