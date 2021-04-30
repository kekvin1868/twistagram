package routing

import (
	"twistagram/src/modules/share/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/postShare", controller.PostShare)
	router.GET("/loadShare/:ID", controller.LoadShare)
}
