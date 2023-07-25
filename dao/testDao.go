package dao

import (
	"testing-gin/domain"
	"testing-gin/model"
)

type TestDao interface {
	Save(person model.Person) (uint, error)
	FindById(personId uint) (model.Person, error)
	Delete(personId uint) error
	SaveORM(person domain.Person) (uint, error)
	FindByIdORM(personId uint) (domain.Person, error)
	DeleteORM(personId uint) error
}
