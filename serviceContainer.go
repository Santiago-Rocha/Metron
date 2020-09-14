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
}

type kernel struct{}

func (k *kernel) InjectTransactionController() api.TransactionHandler {

	mongoDataStore := db.NewDatastore()
	transactionRepository := &repository.TransactionRepository{mongoDataStore}
	transactionService := &service.TransactionService{transactionRepository}
	transactionHandler := api.TransactionHandler{transactionService}

	return transactionHandler
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
