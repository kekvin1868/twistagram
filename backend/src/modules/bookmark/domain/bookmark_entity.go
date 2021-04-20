package domain

import (
	"fmt"
	"time"
	post "twistagram/src/modules/post/domain"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Bookmark{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	orm.Engine.Model(&Bookmark{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Bookmark{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

type Bookmark struct {
	ID        uint      `json:"id" gorm:"'primaryKey; not Null;autoIncrement'"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_Id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	Post      post.Post `gorm:"foreignKey:PostID"`
	User      user.User `gorm:"foreignKey:UserID"`
}
