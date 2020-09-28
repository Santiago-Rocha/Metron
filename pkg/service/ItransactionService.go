package service

import (
	"../data"
)

type ItransactionService interface {
	InsertTransaction(data.Transaction) error
	FindTransaction(id string) (data.Transaction, error)
}
