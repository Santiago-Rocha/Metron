package data

import "time"

type transactionType string

const (
	DEBIT  = transactionType("DEBIT")
	CREDIT = transactionType("CREDIT")
)

type Transaction struct {
	cash            int
	date            time.Time
	description     string
	transactionType *transactionType
}

type Category struct {
	name         string
	subCateories []*Category
}
