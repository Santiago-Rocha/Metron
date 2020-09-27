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
	InjectCategoryController() api.CategoryHandler
}

type kernel struct{}

func (k *kernel) InjectTransactionController() api.TransactionHandler {

	mongoDataStore := db.NewDatastore()
	transactionRepository := &repository.TransactionRepository{MongoDatastore: mongoDataStore}
	transactionService := &service.TransactionService{ItransactionRepository: transactionRepository}
	transactionHandler := api.TransactionHandler{ItransactionService: transactionService}

	return transactionHandler
}

func (k *kernel) InjectCategoryController() api.CategoryHandler {

	mongoDataStore := db.NewDatastore()
	categoryRepository := &repository.CategoryRepository{MongoDatastore: mongoDataStore}
	categoryService := &service.CategoryService{IcategoryRepository: categoryRepository}
	categoryHandler := api.CategoryHandler{IcategoryService: categoryService}

	return categoryHandler
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
