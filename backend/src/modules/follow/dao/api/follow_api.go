package api

type Followers struct {
	ID     uint
	UserID uint `json:"user_id"`
}

type Following struct {
	ID       uint
	FollowID uint `json:"follow_id"`
}

type FollowersRes struct {
	Followers []Followers
	Count     uint64
}

type FollowingRes struct {
	Following []Following
	Count     uint64
}
