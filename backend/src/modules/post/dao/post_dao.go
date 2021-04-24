package dao

import (
	"fmt"
	commentDAO "twistagram/src/modules/comment/dao"
	likeDAO "twistagram/src/modules/like/dao"
	"twistagram/src/modules/post/domain"
	"twistagram/src/modules/post/domain/api"
	"twistagram/src/orm"
)

func GetPost(ID uint64) (*api.PostAPI, error) {
	var post api.PostAPI

	res := orm.Engine.Table("posts").Select("posts.caption, posts.id, posts.user_id,users.full_name").Joins("JOIN users on posts.user_id = users.id").Where("posts.id = ?", ID).First(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	like, _ := likeDAO.GetLikes(ID)
	comment, _ := commentDAO.GetComment(ID)
	post.Like = *like
	post.Comment = *comment
	return &post, nil

}

func GetAllUserPost(UserID uint64) (*[]api.PostAPI, error) {

	var ids []api.PostID
	var posts []api.PostAPI

	res := orm.Engine.Table("posts").Select("posts.id").Where("posts.user_id = ?", UserID).Scan(&ids)

	if res.Error != nil {
		return nil, res.Error
	}

	for val := range ids {
		fmt.Println(val)
	}

	return &posts, nil

}

func LoadFollowingPost(UserID uint64) (*[]api.PostID, error) {
	type FollowingID struct {
		FollowID uint
	}

	var postID []api.PostID
	var follow []FollowingID
	// res := orm.Engine.Table("follows").Select("follows.id,follows.follow_id").Where("user_id = ?", UserID).Scan(&following)
	res := orm.Engine.Table("follows").Select("follows.follow_id").Where("user_id = ?", UserID).Scan(&follow)

	if res.Error != nil {
		return nil, res.Error
	}

	// []arrayID :=
	res = orm.Engine.Table("posts").Select("posts.id").Joins("JOIN users on posts.user_id = users.id").Where("posts.users_id IN", follow).Find(&postID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &postID, nil

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
