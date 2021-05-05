package service

import (
	"twistagram/src/modules/user/dao"
	"twistagram/src/modules/user/domain"
	"twistagram/src/modules/user/domain/api"
)

func GetUserData(ID uint64) (*domain.User, error) {
	return dao.GetUserData(ID)
}

func EditUserData(user *domain.User) (*domain.User, error) {
	return dao.EditUserData(user)
}

func SearchUser(keyword string) (*[]api.SearchAPI, error) {
	return dao.SearchUser(keyword)
}

func GetSugestion(ID uint64) (*[]api.SearchAPI, error) {
	return dao.GetSugestion(ID)
}