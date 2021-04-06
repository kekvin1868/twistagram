package dao

import (
	"twistagram/src/modules/like/domain"
	"twistagram/src/modules/like/domain/api"
	"twistagram/src/orm"
)

func GetLikes(PostID uint64) (*[]api.LikeApi, error) {
	var likes []api.LikeApi

	res := orm.Engine.Table("likes").Select("users.full_name,users.id").Joins("JOIN posts on posts.id = likes.post_id").Joins("JOIN users on likes.user_id = users.id").Scan(&likes)
	if res.Error != nil {
		return nil, res.Error
	}

	return &likes, nil
}

func PostLike(like *domain.Like) (*domain.Like, error) {
	var newLike domain.Like
	newLike = *like

	res := orm.Engine.Create(&newLike)

	if res.Error != nil {
		return nil, res.Error
	}

	return like, nil

}
