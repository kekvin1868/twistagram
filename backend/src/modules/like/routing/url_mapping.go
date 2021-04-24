package routing

import (
	"twistagram/src/modules/like/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.GET("/getLikes/:post_id", controller.GetLikes)
	router.POST("/postLike", controller.PostLike)
}
