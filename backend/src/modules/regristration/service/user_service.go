package service

import (
	"twistagram/src/modules/regristration/dao"
	"twistagram/src/modules/user/domain"
)

func Login(email string, pass string) (*domain.User, error) {
	return dao.Login(email, pass)
}

func Register(user *domain.User) (*domain.User, error) {
	return dao.Register(user)
}
