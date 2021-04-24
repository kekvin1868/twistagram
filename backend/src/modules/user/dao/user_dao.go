package dao

import (
	"twistagram/src/modules/user/domain"
	"twistagram/src/modules/user/domain/api"
	"twistagram/src/orm"
)

func GetUserData(ID uint64) (*domain.User, error) {
	var user domain.User

	res := orm.Engine.First(&user, ID)
	// SELECT * FROM users WHERE id = 'id' LIMIT 1;
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil

}

func SearchUser(keyword string) (*[]api.SearchAPI, error) {
	var user []api.SearchAPI
	keyword = keyword + "%"
	res := orm.Engine.Table("users").Select("users.id,users.full_name").Where("full_name LIKE ?", keyword).Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func EditUserData(user *domain.User) (*domain.User, error) {
	var newUser = user

	res := orm.Engine.Model(&user).Updates(&newUser)

	if res.Error != nil {
		return nil, res.Error
	}

	return newUser, nil

}