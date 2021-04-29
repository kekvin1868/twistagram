package domain

import (
	"fmt"
	"time"
	post "twistagram/src/modules/post/domain"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Share{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}

	orm.Engine.Model(&Share{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Share{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

}

type Share struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	UserID    uint      `json:"user_id"`
	User      user.User `gorm:"foreignKey:UserID"`
	PostID    uint      `json:"post_id"`
	Post      post.Post `gorm:"foreignKey:UserID"`
}
