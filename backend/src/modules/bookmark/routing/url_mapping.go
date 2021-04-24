package routing

import (
	"twistagram/src/modules/bookmark/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/postBookmark", controller.PostBookmark)
	router.GET("/getBookmark/:ID", controller.GetBookmark)
	router.DELETE("/deleteBookmark/:ID", controller.DeleteBookmark)
}
