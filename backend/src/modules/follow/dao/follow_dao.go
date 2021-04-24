package dao

import (
	"twistagram/src/modules/follow/dao/api"
	"twistagram/src/modules/follow/domain"
	"twistagram/src/orm"
)

func isAlreadyFollow(UserID uint64, FollowID uint64) bool {
	var follow domain.Follow
	var followed bool

	if err := orm.Engine.Where("user_id = ? AND follow_id = ?", UserID, FollowID).First(&follow).Error; err != nil {
		followed = false
	} else {
		followed = true
	}

	return followed
}

func PostFollow(follow *domain.Follow) (*domain.Follow, error) {
	var newFollow domain.Follow
	newFollow = *follow

	alreadyFollow := isAlreadyFollow(uint64(newFollow.UserID), uint64(newFollow.FollowID))
	if alreadyFollow == true {
		return nil, nil
	}

	res := orm.Engine.Create(&newFollow)

	if res.Error != nil {
		return nil, res.Error
	}

	return follow, nil
}

func GetFollowers(UserID uint64) (*[]api.FollowAPI, error) {
	var followers []api.FollowAPI

	res := orm.Engine.Table("follows").Select("follows.id,follows.follow_id").Where("user_id = ?", UserID).Scan(&followers)
	if res.Error != nil {
		return nil, res.Error
	}

	return &followers, nil
}

func DeleteFollow(ID uint64) error {
	res := orm.Engine.Delete(&domain.Follow{}, ID)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
