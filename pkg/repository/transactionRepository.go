package repository

import (
	"context"

	"../data"
	"../db"
	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	*db.MongoDatastore
}

func (repository *TransactionRepository) Insert(transaction data.Transaction) error {
	insertedResult, err := repository.GetCollection("transaction").InsertOne(context.TODO(), transaction)
	if err == nil {
		logrus.Infoln("Transaction", insertedResult.InsertedID, "inserted in database")

	}
	return err
}

func (repository *TransactionRepository) GetTransaction(id string) (data.Transaction, error) {
	var transaction data.Transaction
	err := repository.GetCollection("transaction").FindOne(context.TODO(), nil).Decode(&transaction)
	return transaction, err
}
