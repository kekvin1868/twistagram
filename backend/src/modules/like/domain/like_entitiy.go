package domain

import (
	"fmt"
	"time"
	post "twistagram/src/modules/post/domain"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Like{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	orm.Engine.Model(&Like{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Like{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

type Like struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	UserID    uint      `json:"user_id"`
	User      user.User ` gorm:"references:id;foreignKey:UserID"`
	PostID    uint      `json:"post_id"`
	Post      post.Post `gorm:"references:id;foreignKey:PostID"`
}
