package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	Engine   *gorm.DB
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbName   = "twistagram"
)

func init() {
	var err error
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	Engine, err = gorm.Open("postgres", psqlinfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
