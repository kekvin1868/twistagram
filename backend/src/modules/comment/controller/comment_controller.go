package controller

import (
	"net/http"
	"twistagram/src/modules/comment/domain"
	"twistagram/src/modules/comment/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	var comment domain.Comment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newComment *domain.Comment
	newComment, err = service.Post(&comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newComment,
	}

	c.JSON(http.StatusOK, res)
}

func Test(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("post_id"))
	comment, err := service.GetComment(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *comment,
	}

	c.JSON(http.StatusOK, res)

}
