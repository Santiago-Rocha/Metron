package service

import (
	"../data"
)

type ItransactionService interface {
	InsertTransaction(data.Transaction)
	GetTransaction(id string) (data.Transaction, error)
}
