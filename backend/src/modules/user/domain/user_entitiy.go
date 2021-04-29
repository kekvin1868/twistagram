package domain

import (
	"fmt"
	"time"
	"twistagram/src/orm"
)

func init() {
	err := orm.Engine.AutoMigrate(&User{})

	if err != nil {
		fmt.Println(err.Error)
	}
}

type User struct {
	ID        uint      `json:"id" gorm:"'primaryKey;not Null;autoIncrement'"`
	FullName  string    `json:"fullname" gorm:"type:varchar(255);not Null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not Null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not Null"`
	Gender    string    `json:"gender" gorm:"type:varchar(255);not Null"`
	Phone     string    `json:"phone" gorm:"type:varchar(255);not Null"`
	Bio       string    `json:"bio" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated _at" gorm:"autoUpdateTime:milli"`
	Profile   string    `json:"profile" gorm:"type:text"`
}
