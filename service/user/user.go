package user

import (
	"bee-boilerplate/public/datalayer/model"
)

type User struct {
	model.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Password string
}
