package repository

import (
	"../data"
)

type IcategoryRepository interface {
	Insert(data.Category) error
	Find(id string) (data.Category, error)
}
