package dao

import (
	"fmt"
	"twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func isEmailExist(email string) bool {
	var user domain.User

	res := orm.Engine.Where("email = ?", email, &user)
	fmt.Print(user.Bio)
	if res.Error != nil {
		return false
	}

	return true
}

func Register(user *domain.User) (*domain.User, error) {
	var newUser domain.User

	newUser = *user

	emailExist := isEmailExist(newUser.Email)
	fmt.Print(emailExist, newUser.Email)
	if emailExist == true {
		return nil, nil
	}
	res := orm.Engine.Create(&newUser)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil

}

func Login(email string, pass string) (*domain.User, error) {
	var user domain.User

	res := orm.Engine.Where("email = ? AND password = ?", email, pass).Find(&user)
	fmt.Print(email, pass)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
