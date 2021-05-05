package domain

import (
	"fmt"
	"time"
	post "twistagram/src/modules/post/domain"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Report{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	orm.Engine.Model(&Report{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Report{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

type Report struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	UserID    uint      `json:"user_id"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	Post      post.Post `gorm:"foreignKey:PostID"`
	User      user.User `gorm:"foreignKey:UserID"`
}
