package controller

import (
	"net/http"
	"twistagram/src/modules/report/domain"
	"twistagram/src/modules/report/service"
	utils "twistagram/src/utils/Response"

	"github.com/gin-gonic/gin"
)

func PostReport(c *gin.Context) {
	var report domain.Report

	err := c.ShouldBindJSON(&report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newReport *domain.Report
	newReport, err = service.PostReport(&report)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if newReport == nil {
		failed := utils.Response{
			Status:  http.StatusConflict,
			Message: "Already reported",
			Data:    nil,
		}
		c.JSON(http.StatusConflict, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newReport,
	}
	c.JSON(http.StatusOK, res)
}
