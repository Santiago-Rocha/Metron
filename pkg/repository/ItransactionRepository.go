package repository

import (
	"../data"
)

type ItransactionRepository interface {
	Insert(data.Transaction) error
	Find(id string) (data.Transaction, error)
}
