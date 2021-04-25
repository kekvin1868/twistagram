package routing

import (
	"twistagram/src/modules/photo/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/storePhoto", controller.StorePhoto)
	router.GET("/retreivePhoto/:postid", controller.RetreivePhoto)
}
