package dao

import (
	"twistagram/src/modules/photo/domain"
	"twistagram/src/modules/photo/domain/api"
	"twistagram/src/orm"
)

func StorePhoto(photo *domain.Photo) (*domain.Photo, error) {
	var newPhoto domain.Photo

	newPhoto = *photo
	res := orm.Engine.Create(&newPhoto)

	if res.Error != nil {
		return nil, res.Error
	}

	return photo, nil
}

func RetreivePhoto(PostID uint64) (*[]api.PhotoRes, error) {

	var photos []api.PhotoRes
	res := orm.Engine.Table("photos").Select("id,index,encode(content,'base64')").Where("photos.post_id = ?", PostID).Find(&photos)

	if res.Error != nil {
		return nil, res.Error
	}

	return &photos, nil

}
