package dao

import (
	"twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func GetUserData(ID uint64) (*domain.User, error) {
	var user domain.User

	res := orm.Engine.First(&user, ID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil

}

func EditUserData(user *domain.User) (*domain.User, error) {
	var newUser = user

	res := orm.Engine.Model(&user).Updates(&newUser)

	if res.Error != nil {
		return nil, res.Error
	}

	return newUser, nil

}
