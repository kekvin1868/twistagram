package service

import (
	"twistagram/src/modules/follow/dao"
	"twistagram/src/modules/follow/dao/api"
	"twistagram/src/modules/follow/domain"
)

func PostFollow(follow *domain.Follow) (*domain.Follow, error) {
	return dao.PostFollow(follow)
}

func GetFollowers(UserID uint64) (*api.FollowersRes, error) {
	return dao.GetFollowers(UserID)
}

func DeleteFollow(ID uint64) error {
	return dao.DeleteFollow(ID)
}

func GetFollowing(UserID uint64) (*api.FollowingRes, error) {
	return dao.GetFollowing(UserID)
}
