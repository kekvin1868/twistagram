package routing

import (
	"twistagram/src/modules/post/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.GET("/getPost/:id", controller.GetPost)
	router.GET("getAllUserPost/:uid", controller.GetAllUserPost)
	router.POST("/post", controller.Post)
	router.GET("/loadFeeds/:uid", controller.LoadFollowingPost)
	router.PATCH("/editPost", controller.EditPost)
}
