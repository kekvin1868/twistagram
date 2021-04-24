package controller

import (
	"net/http"
	"twistagram/src/modules/post/domain"
	"twistagram/src/modules/post/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("id"))
	post, err := service.GetPost(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *post,
	}

	c.JSON(http.StatusOK, res)

}

func GetAllUserPost(c *gin.Context) {
	UserID, _ := parser.ParseID(c.Param("uid"))
	posts, err := service.GetAllUserPost(UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *posts,
	}

	c.JSON(http.StatusOK, res)

}

func Post(c *gin.Context) {
	var post domain.Post

	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newPost *domain.Post
	newPost, err = service.Post(&post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status:  http.StatusOK,
		Message: "posted",
		Data:    *newPost,
	}

	c.JSON(http.StatusOK, res)

}
