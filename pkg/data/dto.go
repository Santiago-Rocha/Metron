package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionDTO struct {
	Cash            int                `json:"cash" bson:"cash"`
	Date            primitive.DateTime `json:"date" bson:"date"`
	Description     string             `json:"description" bson:"description"`
	TransactionType TransactionType    `json:"transactionType" bson:"transactionType"`
	Category        CategoryDTO        `json:"category" bson:"category"`
}

type CategoryDTO struct {
	Name          string            `json:"name" bson:"name"`
	SubCategories map[string]string `json:"SubCategories" bson:"SubCategories"`
}
