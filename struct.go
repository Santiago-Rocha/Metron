package main

import "time"

type transactionType string

const (
	DEBIT  = transactionType("DEBIT")
	CREDIT = transactionType("CREDIT")
)

type transaction struct {
	cash            int
	date            time.Time
	description     string
	transactionType *transactionType
}

type category struct {
	name         string
	subCateories []*category
}
