package service

import (
	"twistagram/src/modules/comment/dao"
	"twistagram/src/modules/comment/domain"
	dto "twistagram/src/modules/comment/domain/api"
)

func Post(comment *domain.Comment) (*domain.Comment, error) {
	return dao.Post(comment)
}

func GetComment(PostID uint64) (*[]dto.CommentAPI, error) {
	return dao.GetComment(PostID)
}
