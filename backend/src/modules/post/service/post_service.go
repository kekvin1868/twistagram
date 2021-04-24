package service

import (
	"twistagram/src/modules/post/dao"
	"twistagram/src/modules/post/domain"
	"twistagram/src/modules/post/domain/api"
)

func GetPost(ID uint64) (*api.PostAPI, error) {
	return dao.GetPost(ID)
}

func Post(post *domain.Post) (*domain.Post, error) {
	return dao.Post(post)
}

func GetAllUserPost(UserID uint64) (*[]api.PostRes, error) {
	return dao.GetAllUserPost(UserID)
}

func LoadFollowingPost(UserID uint64) (*[]api.PostID, error) {
	return dao.LoadFollowingPost(UserID)
}
