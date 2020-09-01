package repository

import (
	"context"

	"../data"
	"../db"
)

type transactionRepository struct {
	dataStore *db.MongoDatastore
}

func (repository *transactionRepository) InsertTransaction(transaction data.Transaction) {
	repository.createDS()
	repository.dataStore.GetCollection("transaction").InsertOne(context.TODO(), transaction)
}

func (repository *transactionRepository) GetTransaction(id string) (data.Transaction, error) {
	repository.createDS()
	var transaction data.Transaction
	err := repository.dataStore.GetCollection("transaction").FindOne(context.TODO(), nil).Decode(&transaction)
	return transaction, err
}

func (repository *transactionRepository) createDS() {
	if repository.dataStore == nil {
		repository.dataStore = db.NewDatastore()
	}
}
