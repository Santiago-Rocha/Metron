package repository

import (
	"context"

	"../data"
	"../db"
	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	*db.MongoDatastore
}

func (repository *CategoryRepository) Insert(category data.Category) error {
	insertedResult, err := repository.GetCollection("category").InsertOne(context.TODO(), category)
	if err == nil {
		logrus.Infoln("Category", insertedResult.InsertedID, "inserted in database")

	}
	return err
}

func (repository *CategoryRepository) Find(id string) (data.Category, error) {
	var category data.Category
	err := repository.GetCollection("category").FindOne(context.TODO(), nil).Decode(&category)
	return category, err
}
