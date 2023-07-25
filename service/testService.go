package service

import (
	"testing-gin/domain"
	"testing-gin/model"
)

type TestService interface {
	CalculateFibonacci(number int64) int64
	CreateGetDeletePersonTestCase(person model.Person) (model.Person, error)
	CreateGetDeletePersonORMTestCase(person domain.Person) (domain.Person, error)
}
