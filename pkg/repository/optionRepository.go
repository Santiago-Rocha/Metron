package repository

import (
	"context"

	"../data"
	"../db"
	"github.com/sirupsen/logrus"
)

type OptionRepository struct {
	*db.MongoDatastore
}

func (repository *OptionRepository) Insert(option data.Option) error {
	insertedResult, err := repository.GetCollection("option").InsertOne(context.TODO(), option)
	if err == nil {
		logrus.Infoln("Option", insertedResult.InsertedID, "inserted in database")

	}
	return err
}

func (repository *OptionRepository) Find(id string) (data.Option, error) {
	var option data.Option
	err := repository.GetCollection("option").FindOne(context.TODO(), nil).Decode(&option)
	return option, err
}
