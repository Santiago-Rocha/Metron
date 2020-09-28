package service

import (
	"../data"
	"../repository"
)

type CategoryService struct {
	repository.IcategoryRepository
}

func (service *CategoryService) InsertCategory(category data.Category) error {
	return service.Insert(category)
}

func (service *CategoryService) FindCategory(id string) (data.Category, error) {
	return service.Find(id)
}
