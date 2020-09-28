package service

import (
	"../data"
	"../repository"
)

type OptionService struct {
	repository.IoptionRepository
}

func (service *OptionService) InsertOption(option data.Option) error {
	return service.Insert(option)
}

func (service *OptionService) FindOption(id string) (data.Option, error) {
	return service.FindOption(id)
}
