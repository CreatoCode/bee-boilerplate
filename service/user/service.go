package user

import (
	"bee-boilerplate/public/datalayer"
	"bee-boilerplate/public/datalayer/model"
)

type User struct {
	model.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	ThirdId  string `json:"thirdId"`
}

func (u *User) AddNew() bool {
	err := datalayer.Create(*u)
	return err != nil
}

func (u *User) CreateIfNotExist() bool {
	r, err := datalayer.IsExist(*u, "uuid=?", u.UUID)
	if !r {
		return u.AddNew()
	}
	// err := datalayer.Create(*u)
	return err != nil
}
