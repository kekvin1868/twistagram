package controller

import (
	"net/http"
	"twistagram/src/modules/like/domain"
	"twistagram/src/modules/like/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func GetLikes(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("post_id"))
	likes, err := service.GetLikes(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *likes,
	}

	c.JSON(http.StatusOK, res)

}

func PostLike(c *gin.Context) {
	var like domain.Like

	err := c.ShouldBindJSON(&like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newLike *domain.Like
	newLike, err = service.PostLike(&like)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if newLike == nil {
		failed := utils.Response{
			Status:  http.StatusConflict,
			Message: "already liked",
			Data:    nil,
		}
		c.JSON(http.StatusConflict, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newLike,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteLike(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	err := service.DeleteLike(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
