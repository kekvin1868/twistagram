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

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newLike,
	}

	c.JSON(http.StatusOK, res)
}
