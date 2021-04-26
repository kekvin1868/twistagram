package dao

import (
	commentDAO "twistagram/src/modules/comment/dao"
	likeDAO "twistagram/src/modules/like/dao"
	photoDAO "twistagram/src/modules/photo/dao"
	"twistagram/src/modules/post/domain"
	"twistagram/src/modules/post/domain/api"
	"twistagram/src/orm"
	// "github.com/lib/pq"
)

func GetPost(ID uint64) (*api.PostAPI, error) {
	var post api.PostAPI

	res := orm.Engine.Table("posts").Select("posts.caption, posts.id, posts.user_id,users.full_name").Joins("JOIN users on posts.user_id = users.id").Where("posts.id = ?", ID).First(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	like, _ := likeDAO.GetLikes(ID)
	comment, _ := commentDAO.GetComment(ID)
	photo, _ := photoDAO.RetreivePhoto(ID)
	post.Like = *like
	post.Comment = *comment
	post.Photo = *photo
	return &post, nil

}

func GetAllUserPost(UserID uint64) (*[]api.PostRes, error) {

	var post []api.PostRes
	res := orm.Engine.Table("posts").Select("posts.id,users.full_name").Joins("JOIN users on posts.user_id = users.id").Where("user_id = ?", UserID).Find(&post)

	if res.Error != nil {
		return nil, res.Error
	}

	return &post, nil
}

func LoadFollowingPost(UserID uint64) (*[]api.PostID, error) {
	type FollowingID struct {
		FollowID uint
	}

	var postID []api.PostID
	// follow := []int64{2, 3}

	// res := orm.Engine.Table("follows").Select("follows.id,follows.follow_id").Where("user_id = ?", UserID).Scan(&following)
	// res := orm.Engine.Table("follows").Select("follows.follow_id").Where("user_id = ?", UserID).Scan(&follow)

	res := orm.Engine.Raw("SELECT posts.id FROM posts WHERE posts.user_id IN (SELECT follows.follow_id FROM follows WHERE follows.user_id = ?)", UserID).Scan(&postID)

	if res.Error != nil {
		return nil, res.Error
	}

	// fmt.Println(follow)

	// res = orm.Engine.Table("posts").Select("posts.id").Where("user_id IN ?", []int64{2, 3}).Find(&postID)

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

func EditPost(post *domain.Post) (*domain.Post, error) {
	var editedPost = post

	res := orm.Engine.Model(&post).Update(&editedPost)

	if res.Error != nil {
		return nil, res.Error
	}

	return editedPost, nil
}
