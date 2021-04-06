package domain

import (
	"fmt"
	"time"
	"twistagram/src/modules/user/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Post{})
	err = orm.Engine.Model(&Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	if err != nil {
		fmt.Println(err.Error)
	}

}

type Post struct {
	ID        uint        `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	Caption   string      `json:"caption" gorm:"type:varchar(255)"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time   `json:"updated _at" gorm:"autoUpdateTime:milli"`
	UserID    uint        `json:"user_id"`
	User      domain.User ` gorm:"references:id;foreignKey:UserID;PRELOAD:true"`
}
