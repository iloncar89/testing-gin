package domain

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Id          uint   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	YearOfBirth int    `json:"yearOfBirth"`
}
