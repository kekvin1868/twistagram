package domain

import (
	"fmt"
	"time"
	post "twistagram/src/modules/post/domain"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Comment{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	orm.Engine.Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

}

type Comment struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content" gorm:"'type:varchar(255)'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	User      user.User `gorm:"foreignKey:UserID"`
	Post      post.Post `gorm:"foreignKey:PostID"`
}
