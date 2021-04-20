package service

import (
	"twistagram/src/modules/like/dao"
	"twistagram/src/modules/like/domain"
	"twistagram/src/modules/like/domain/api"
)

func GetLikes(PostID uint64) (*[]api.LikeApi, error) {
	return dao.GetLikes(PostID)
}

func PostLike(like *domain.Like) (*domain.Like, error) {
	return dao.PostLike(like)
}

func DeleteLike(ID uint64) error {
	return dao.DeleteLike(ID)
}
