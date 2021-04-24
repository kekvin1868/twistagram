package api

type BookmarkAPI struct {
	ID     uint64
	PostID uint64 `json:"post_id"`
}
