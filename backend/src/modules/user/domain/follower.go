package domain

type Follower struct {
	id       int
	idUser   User
	toIDUser User
}
