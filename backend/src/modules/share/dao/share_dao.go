package dao

import (
	"twistagram/src/modules/share/domain"
	"twistagram/src/modules/share/domain/api"
	"twistagram/src/orm"
)

func isShared(UserID uint64, PostID uint64) bool {
	var share domain.Share
	var exist bool

	if err := orm.Engine.Where("post_id = ? AND user_id = ?", PostID, UserID).First(&share).Error; err != nil {
		exist = false
	} else {
		exist = true
	}
	return exist
}

func PostShare(share *domain.Share) (*domain.Share, error) {
	var newShare domain.Share
	newShare = *share

	shareExist := isShared(uint64(newShare.UserID), uint64(newShare.PostID))
	if shareExist == true {
		return nil, nil
	}

	res := orm.Engine.Create(&newShare)

	if res.Error != nil {
		return nil, res.Error
	}

	return share, nil

}

func LoadShare(UserID uint64) (*[]api.ShareAPI, error) {
	var shares []api.ShareAPI

	res := orm.Engine.Table("shares").Raw("SELECT shares.id, shares.user_id, shares.post_id WHERE shares.user_id IN (SELECT follows.follow_id from follows WHERE follows.user_id = ?)", UserID).Find(&shares)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return &shares, nil
}
