package routing

import (
	"github.com/gin-gonic/gin"

	"twistagram/src/modules/follow/controller"
)

func MapURLs(router *gin.Engine) {
	router.POST("/postFollow", controller.PostFollow)
	router.GET("/getFollowers/:ID", controller.GetFollowers)
	router.DELETE("/deleteFollow/:ID", controller.DeleteBookmark)
	router.GET("/getFollowing/:ID", controller.GetFollowing)
}
