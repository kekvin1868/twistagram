package routing

import (
	"twistagram/src/modules/regristration/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
}
