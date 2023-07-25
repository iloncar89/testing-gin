package service

import (
	"testing-gin/dao"
	"testing-gin/domain"
	"testing-gin/model"
)

type TestServiceImpl struct {
	TestDao dao.TestDao
}

func NewTestServiceImpl(testDao dao.TestDao) TestService {
	return &TestServiceImpl{TestDao: testDao}
}

func (t *TestServiceImpl) CalculateFibonacci(number int64) int64 {
	if number <= 1 {
		return number
	}
	return t.CalculateFibonacci(number-1) + t.CalculateFibonacci(number-2)
}

func (t *TestServiceImpl) CreateGetDeletePersonTestCase(person model.Person) (model.Person, error) {
	createdPersonId, err := t.TestDao.Save(person)
	if err != nil {
		return person, err
	}
	findResult, err := t.TestDao.FindById(createdPersonId)
	if err != nil {
		return person, err
	}
	err = t.TestDao.Delete(createdPersonId)
	if err != nil {
		return person, err
	}

	return findResult, nil
}

func (t *TestServiceImpl) CreateGetDeletePersonORMTestCase(person domain.Person) (domain.Person, error) {
	createdPersonId, err := t.TestDao.SaveORM(person)
	if err != nil {
		return person, err
	}
	findResult, err := t.TestDao.FindByIdORM(createdPersonId)
	if err != nil {
		return person, err
	}
	err = t.TestDao.DeleteORM(createdPersonId)
	if err != nil {
		return person, err
	}

	return findResult, nil
}
