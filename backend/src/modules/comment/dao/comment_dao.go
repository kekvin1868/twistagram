package dao

import (
	"twistagram/src/modules/comment/domain"
	dto "twistagram/src/modules/comment/domain/api"
	"twistagram/src/orm"
)

func Post(comment *domain.Comment) (*domain.Comment, error) {
	var newComment domain.Comment
	newComment = *comment

	res := orm.Engine.Create(&newComment)

	if res.Error != nil {
		return nil, res.Error
	}

	return comment, nil

}

func GetComment(PostID uint64) (*[]dto.CommentAPI, error) {
	var comment []dto.CommentAPI

	// res := orm.Engine.Model(&user.User{}).Joins("left join comments on comments.user_id = users.id").Scan(&comment)
	res := orm.Engine.Table("comments").Select("users.full_name,comments.content").Joins("JOIN posts on comments.post_id = posts.id").Joins("JOIN users on comments.user_id = users.id").Where("comments.post_id = ?", PostID).Scan(&comment)
	if res.Error != nil {
		return nil, res.Error
	}

	return &comment, nil
}
