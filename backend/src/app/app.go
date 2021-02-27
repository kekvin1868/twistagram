package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func init() {
	router, _ = generateEngine()
}

func StartApplication() {
	MapURLs()
	err := router.Run(":8081")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func generateEngine() (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default())
	return router, nil
}
