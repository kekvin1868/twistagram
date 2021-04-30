package controller

import (
	"net/http"
	"twistagram/src/modules/share/domain"
	"twistagram/src/modules/share/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func PostShare(c *gin.Context) {
	var share domain.Share

	err := c.ShouldBindJSON(&share)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newShare *domain.Share
	newShare, err = service.PostShare(&share)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if newShare == nil {
		failed := utils.Response{
			Status:  http.StatusConflict,
			Message: "Already Shared",
			Data:    nil,
		}
		c.JSON(http.StatusConflict, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newShare,
	}

	c.JSON(http.StatusOK, res)
}

func LoadShare(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	share, err := service.LoadShare(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if share == nil {
		failed := utils.Response{
			Status:  http.StatusNoContent,
			Message: "No content found",
			Data:    nil,
		}
		c.JSON(http.StatusNoContent, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *share,
	}

	c.JSON(http.StatusOK, res)
}
