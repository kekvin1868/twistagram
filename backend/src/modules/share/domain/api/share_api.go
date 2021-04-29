package api

import post "twistagram/src/modules/post/domain/api"

type ShareAPI struct {
	ID     uint64
	UserID uint64
	Posts  []post.PostAPI
}
