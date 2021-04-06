package routing

import (
	"twistagram/src/modules/comment/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/postComment", controller.Post)
	router.GET("/test/:post_id", controller.Test)
}
