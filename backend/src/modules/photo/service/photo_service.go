package service

import (
	"twistagram/src/modules/photo/dao"
	"twistagram/src/modules/photo/domain"
	"twistagram/src/modules/photo/domain/api"
)

func StorePhoto(post *domain.Photo) (*domain.Photo, error) {
	return dao.StorePhoto(post)
}

func RetreivePhoto(PostID uint64) (*[]api.PhotoRes, error) {
	return dao.RetreivePhoto(PostID)
}
