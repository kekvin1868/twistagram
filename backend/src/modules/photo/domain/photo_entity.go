package domain

import (
	"fmt"
	"time"
	"twistagram/src/modules/post/domain"
	"twistagram/src/orm"

	"github.com/jinzhu/gorm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Photo{})
	orm.Engine.Model(&Photo{}).AddForeignKey("psot_id", "posts(id)", "RESTRICT", "RESTRICT")

	if err != nil {
		fmt.Println(err.Error)
	}

}

type Photo struct {
	gorm.Model
	Index     int           `json:"index" gorm:"'type:int(2);not Null'"`
	Url       string        `json:"url" gorm:"'type:varchar(255);not Null'"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time     `json:"updated _at" gorm:"autoUpdateTime:milli"`
	PostID    uint          `json:"post_id"`
	Post      []domain.Post `gorm:"references:id;foreignKey:PostID"`
}
