package domain

import (
	"fmt"
	"time"
	"twistagram/src/modules/post/domain"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&Photo{})
	orm.Engine.Model(&Photo{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

	if err != nil {
		fmt.Println(err.Error)
	}

}

type Photo struct {
	ID        uint          `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	Index     int           `json:"index" gorm:"'type:int(2);not Null'"`
	Content   []byte        `json:"url" gorm:"'type:bytea;not Null'"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime:milli"`
	PostID    uint          `json:"post_id"`
	Post      []domain.Post `gorm:"references:id;foreignKey:PostID"`
}
