package routing

import (
	"twistagram/src/modules/report/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {
	router.POST("/report", controller.PostReport)
	router.GET("/getReport/:ID", controller.GetReportCounts)
}
