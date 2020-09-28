package repository

import (
	"../data"
)

type IoptionRepository interface {
	Insert(data.Option) error
	Find(id string) (data.Option, error)
}
