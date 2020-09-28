package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionType string

const (
	DEBIT  = transactionType("DEBIT")
	CREDIT = transactionType("CREDIT")
)

type Transaction struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Cash            int                `json:"cash" bson:"cash"`
	Date            primitive.DateTime `json:"date" bson:"date"`
	Description     string             `json:"description" bson:"description"`
	TransactionType *transactionType   `json:"transactionType" bson:"transactionType"`
	Category        primitive.ObjectID `json:"category" bson:"category"`
}

type Category struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	SubCategories map[string]string  `json:"SubCategories" bson:"SubCategories"`
}

type Option struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	OptionsSet []string           `json:"options" bson:"options"`
}
