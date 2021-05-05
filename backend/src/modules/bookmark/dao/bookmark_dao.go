package dao

import (
	"twistagram/src/modules/bookmark/dao/api"
	"twistagram/src/modules/bookmark/domain"
	"twistagram/src/orm"
)

func isBookmarkExist(UserID uint64, PostID uint64) bool {
	var bookmark domain.Bookmark
	var exist bool

	if err := orm.Engine.Where("post_id = ? AND user_id = ?", PostID, UserID).First(&bookmark).Error; err != nil {
		exist = false
	} else {
		exist = true
	}

	return exist

}

func PostBookmark(bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	var newBookmark domain.Bookmark
	newBookmark = *bookmark

	bookmarkExist := isBookmarkExist(uint64(newBookmark.UserID), uint64(newBookmark.PostID))
	if bookmarkExist == true {
		return nil, nil
	}

	res := orm.Engine.Create(&newBookmark)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookmark, nil
}

func GetBookmark(UserID uint64) (*[]api.BookmarkAPI, error) {
	var bookmarks []api.BookmarkAPI

	res := orm.Engine.Table("bookmarks").Select("bookmarks.id,bookmarks.post_id").Joins("JOIN posts on bookmarks.post_id = posts.id").Joins("JOIN users on bookmarks.user_id = users.id").Where("bookmarks.user_id = ?", UserID).Scan(&bookmarks)
	if res.Error != nil {
		return nil, res.Error
	}

	return &bookmarks, nil
}

func DeleteBookmark(ID uint64) error {
	res := orm.Engine.Delete(&domain.Bookmark{}, ID)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
