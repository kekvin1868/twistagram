package api

type LikeApi struct {
	ID       uint
	UserID   uint   `json:"user_id"`
	FullName string `json:"fullname"`
}
