package models

import (
	"pr-prueba-svc/kit/libs"
)

var User *user

type user struct {
	Account  string
	Password string
}

func (this *user) Add(account string, password string) (user, error) {
	var data user
	data.Account = account
	data.Password = password
	if err := libs.DB.Create(&data).Error; err != nil {
		return user{}, err
	} else {

		return data, nil
	}

}
