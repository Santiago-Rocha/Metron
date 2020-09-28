package main

import (
	"sync"

	"./pkg/api"
	"./pkg/db"
	"./pkg/repository"
	"./pkg/service"
)

type IServiceContainer interface {
	InjectTransactionController() api.TransactionHandler
	InjectOptionController() api.OptionHandler
}

type kernel struct{}

func (k *kernel) InjectTransactionController() api.TransactionHandler {

	mongoDataStore := db.NewDatastore()
	transactionRepository := &repository.TransactionRepository{MongoDatastore: mongoDataStore}
	transactionService := &service.TransactionService{ItransactionRepository: transactionRepository}
	transactionHandler := api.TransactionHandler{ItransactionService: transactionService}

	return transactionHandler
}

func (k *kernel) InjectOptionController() api.OptionHandler {

	mongoDataStore := db.NewDatastore()
	optionRepository := &repository.OptionRepository{MongoDatastore: mongoDataStore}
	optionService := &service.OptionService{IoptionRepository: optionRepository}
	optionHandler := api.OptionHandler{IoptionService: optionService}

	return optionHandler
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
