package domain

import (
	"fmt"
	"time"
	user "twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Follow{})
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	orm.Engine.Model(&Follow{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	orm.Engine.Model(&Follow{}).AddForeignKey("follow_id", "users(id)", "RESTRICT", "RESTRICT")
}

type Follow struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	UserID    uint      `json:"user_id"`
	FollowID  uint      `json:"follow_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	User      user.User `gorm:"foreignKey:UserID"`
	Follow    user.User `gorm:"foreginKey:FollowID"`
}
