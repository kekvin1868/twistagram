package dao

import (
	"fmt"
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
	res := orm.Engine.Table("users").Select("users.profile,users.id,users.full_name").Where("full_name LIKE ?", keyword).Find(&user)
	var count int64
	count = res.RowsAffected
	fmt.Println(user)
	if res.Error != nil {
		return nil, res.Error
	}

	if count == 0 {
		return nil, nil
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

func GetSugestion(UserID uint64) (*[]api.SearchAPI, error) {
	var user []api.SearchAPI
	res := orm.Engine.Raw("SELECT id, full_name, profile FROM users WHERE users.id NOT IN(SELECT follows.follow_id FROM follows WHERE follows.user_id = ?) LIMIT 3",UserID).Find(&user)
	
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
