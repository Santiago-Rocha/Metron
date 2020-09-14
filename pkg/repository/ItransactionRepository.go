package repository

import (
	"../data"
)

type ItransactionRepository interface {
	Insert(data.Transaction) error
	GetTransaction(id string) (data.Transaction, error)
}
