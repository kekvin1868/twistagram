package api

type LikeApi struct {
	ID       uint
	UserID   uint   `json:"user_id"`
	Fullname string `json:"fullname"`
}
