package dao

import (
	"twistagram/src/modules/like/domain"
	"twistagram/src/modules/like/domain/api"
	"twistagram/src/orm"
)

func isLikeExist(PostID uint64, UserID uint64) bool {
	var like domain.Like
	var exist bool

	if err := orm.Engine.Where("post_id = ? AND user_id = ?", PostID, UserID).First(&like).Error; err != nil {
		exist = false
	} else {
		exist = true
	}

	return exist
}

func GetLikes(PostID uint64) (*[]api.LikeApi, error) {
	var likes []api.LikeApi

	res := orm.Engine.Table("likes").Select("likes.id , users.full_name , likes.user_id").Joins("JOIN posts on posts.id = likes.post_id").Joins("JOIN users on likes.user_id = users.id").Where("likes.post_id = ?", PostID).Scan(&likes)
	if res.Error != nil {
		return nil, res.Error
	}

	return &likes, nil
}

func PostLike(like *domain.Like) (*domain.Like, error) {
	var newLike domain.Like
	newLike = *like

	likeExist := isLikeExist(uint64(newLike.PostID), uint64(newLike.UserID))
	if likeExist == true {
		return nil, nil
	}

	res := orm.Engine.Create(&newLike)

	if res.Error != nil {
		return nil, res.Error
	}

	return &newLike, nil

}

func DeleteLike(ID uint64) error {
	res := orm.Engine.Delete(&domain.Like{}, ID)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
