package dao

import (
	commentDAO "twistagram/src/modules/comment/dao"
	likeDAO "twistagram/src/modules/like/dao"
	"twistagram/src/modules/post/domain"
	"twistagram/src/modules/post/domain/api"
	"twistagram/src/orm"
)

func GetPost(ID uint64) (*api.PostAPI, error) {
	var post api.PostAPI

	res := orm.Engine.Table("posts").Select("posts.caption, posts.id, users.id,users.full_name").Joins("JOIN users on posts.user_id = users.id").Scan(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	like, _ := likeDAO.GetLikes(ID)
	comment, _ := commentDAO.GetComment(ID)
	post.Like = *like
	post.Comment = *comment
	return &post, nil

}

func GetAllUserPost(UserID uint64) (*[]domain.Post, error) {
	var post []domain.Post

	res := orm.Engine.Where("user_id = ?", UserID).Find(&post)

	if res.Error != nil {
		return nil, res.Error
	}

	return &post, nil

}

func Post(post *domain.Post) (*domain.Post, error) {
	var newPost domain.Post

	newPost = *post
	res := orm.Engine.Create(&newPost)

	if res.Error != nil {
		return nil, res.Error
	}

	return post, nil
}
