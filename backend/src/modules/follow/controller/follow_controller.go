package controller

import (
	"net/http"
	"twistagram/src/modules/follow/domain"
	"twistagram/src/modules/follow/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func PostFollow(c *gin.Context) {
	var follow domain.Follow

	err := c.ShouldBindJSON(&follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newFollow *domain.Follow
	newFollow, err = service.PostFollow(&follow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if newFollow == nil {
		failed := utils.Response{
			Status:  http.StatusConflict,
			Message: "already followed",
			Data:    nil,
		}
		c.JSON(http.StatusConflict, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newFollow,
	}

	c.JSON(http.StatusOK, res)
}

func GetFollowers(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	followers, err := service.GetFollowers(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *followers,
	}

	c.JSON(http.StatusOK, res)
}

func GetFollowing(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	following, err := service.GetFollowing(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *following,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteBookmark(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	err := service.DeleteFollow(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
