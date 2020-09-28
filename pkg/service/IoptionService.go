package service

import (
	"../data"
)

type IoptionService interface {
	InsertOption(data.Option) error
	FindOption(id string) (data.Option, error)
}
