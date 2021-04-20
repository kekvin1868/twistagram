package service

import (
	"twistagram/src/modules/bookmark/dao"
	"twistagram/src/modules/bookmark/dao/api"
	"twistagram/src/modules/bookmark/domain"
)

func PostBookmark(bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	return dao.PostBookmark(bookmark)
}

func GetBookmark(UserID uint64) (*[]api.BookmarkAPI, error) {
	return dao.GetBookmark(UserID)
}

func DeleteBookmark(ID uint64) error {
	return dao.DeleteBookmark(ID)
}
