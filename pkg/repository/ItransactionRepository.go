package repository

import (
	"../data"
)

type ItransactionRepository interface {
	InsertTransaction(data.Transaction)
	GetTransaction(id string) (data.Transaction, error)
}
