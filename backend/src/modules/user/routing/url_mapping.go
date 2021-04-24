package routing

import (
	"twistagram/src/modules/user/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.GET("/getUserData/:id", controller.GetUserData)
	router.PATCH("/updateUserData", controller.EditUserData)
	router.GET("/searchUser/:key", controller.SearchUser)
}
