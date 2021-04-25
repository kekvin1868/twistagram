package api

import (
	comment "twistagram/src/modules/comment/domain/api"
	like "twistagram/src/modules/like/domain/api"
	photo "twistagram/src/modules/photo/domain/api"
)

type PostAPI struct {
	ID       uint                 `json:"id"`
	Caption  string               `json:"caption"`
	FullName string               `json:"fullname"`
	UserID   uint                 `json:"user_id"`
	Like     []like.LikeApi       `json:"like"`
	Comment  []comment.CommentAPI `json:"comment"`
	Photo    []photo.PhotoRes     `json:"photo"`
}

type PostRes struct {
	ID       uint
	FullName string `json:"fullname"`
}

type PostID struct {
	ID uint
}
