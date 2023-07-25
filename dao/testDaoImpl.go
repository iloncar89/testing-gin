package dao

import (
	"database/sql"
	"errors"
	"testing-gin/domain"
	"testing-gin/model"

	"gorm.io/gorm"
)

type TestDaoImpl struct {
	Db    *sql.DB
	DbOrm *gorm.DB
}

func NewPersonDaoImpl(Db *sql.DB, DbOrm *gorm.DB) TestDao {
	return &TestDaoImpl{Db: Db, DbOrm: DbOrm}
}

func (t *TestDaoImpl) Save(person model.Person) (uint, error) {
	statement, err := t.Db.Prepare("INSERT INTO person (first_name, last_name, year_of_birth) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	var lastInsertID uint
	result := statement.QueryRow(person.FirstName, person.LastName, person.YearOfBirth)

	if result.Err() != nil {
		return 0, result.Err()
	}

	if err := result.Scan(&lastInsertID); err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (t *TestDaoImpl) FindById(personId uint) (model.Person, error) {
	var person model.Person
	statement, err := t.Db.Prepare("SELECT id, first_name, last_name, year_of_birth FROM person WHERE id=$1;")
	if err != nil {
		return person, err
	}
	defer statement.Close()

	result := statement.QueryRow(personId)

	if result.Err() != nil {
		return person, result.Err()
	}

	if err := result.Scan(&person.Id, &person.FirstName, &person.LastName, &person.YearOfBirth); err != nil {
		return model.Person{}, err
	}
	return person, nil
}

func (t *TestDaoImpl) Delete(personId uint) error {
	statement, err := t.Db.Prepare("DELETE FROM person WHERE id=$1;")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(personId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TestDaoImpl) SaveORM(person domain.Person) (uint, error) {
	result := t.DbOrm.Table("person").Create(&person)
	if result.Error != nil {
		return 0, result.Error
	}
	return person.Id, nil
}

func (t *TestDaoImpl) FindByIdORM(personId uint) (domain.Person, error) {
	var person domain.Person
	result := t.DbOrm.Table("person").Find(&person, personId)
	if result != nil {
		return person, nil
	} else {
		return person, errors.New("person is not found")
	}
}

func (t *TestDaoImpl) DeleteORM(personId uint) error {
	var person domain.Person
	result := t.DbOrm.Table("person").Where("id = ?", personId).Delete(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
