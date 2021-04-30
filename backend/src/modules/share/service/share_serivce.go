package service

import (
	"twistagram/src/modules/share/dao"
	"twistagram/src/modules/share/domain"
	"twistagram/src/modules/share/domain/api"
)

func PostShare(share *domain.Share) (*domain.Share, error) {
	return dao.PostShare(share)
}

func LoadShare(UserID uint64) (*[]api.ShareAPI, error) {
	return dao.LoadShare(UserID)
}
