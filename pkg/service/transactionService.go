package service

import (
	"../data"
	"../repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionService struct {
	repository.ItransactionRepository
	IcategoryService
}

func (service *TransactionService) InsertTransaction(transactionDTO data.TransactionDTO) error {
	var transaction = data.Transaction{
		Id:              primitive.NewObjectID(),
		Cash:            transactionDTO.Cash,
		Date:            transactionDTO.Date,
		Description:     transactionDTO.Description,
		TransactionType: &transactionDTO.TransactionType,
	}

	var category = data.Category{
		Id:            primitive.NewObjectID(),
		Name:          transactionDTO.Category.Name,
		SubCategories: transactionDTO.Category.SubCategories,
	}

	err := service.InsertCategory(category)
	if err == nil {
		transaction.Category = category.Id
		return service.Insert(transaction)
	}

	return err

}

func (service *TransactionService) FindTransaction(id string) (data.Transaction, error) {
	return service.Find(id)
}
