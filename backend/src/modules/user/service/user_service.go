package service

import (
	"twistagram/src/modules/user/dao"
	"twistagram/src/modules/user/domain"
)

func GetUserData(ID uint64) (*domain.User, error) {
	return dao.GetUserData(ID)
}

func EditUserData(user *domain.User) (*domain.User, error) {
	return dao.EditUserData(user)
}
