package service

import (
	"../data"
)

type IcategoryService interface {
	InsertCategory(data.Category) error
	GetCategory(id string) (data.Category, error)
}
