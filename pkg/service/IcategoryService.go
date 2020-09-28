package service

import (
	"../data"
)

type IcategoryService interface {
	InsertCategory(data.Category) error
	FindCategory(id string) (data.Category, error)
}
