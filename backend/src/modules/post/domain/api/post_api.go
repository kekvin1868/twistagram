package api

import (
	comment "twistagram/src/modules/comment/domain/api"
	like "twistagram/src/modules/like/domain/api"
)

type PostAPI struct {
	ID       uint                 `json:"id"`
	Caption  string               `json:"caption"`
	Fullname string               `json:"fullname"`
	UserID   uint                 `json:"user_id"`
	Like     []like.LikeApi       `json:"like"`
	Comment  []comment.CommentAPI `json:"comment"`
}
