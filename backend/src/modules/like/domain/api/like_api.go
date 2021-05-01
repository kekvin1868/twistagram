package api

type LikeApi struct {
	ID       uint	`json:"id"`
	UserID   uint   `json:"user_id"`
	FullName string `json:"fullname"`
}
