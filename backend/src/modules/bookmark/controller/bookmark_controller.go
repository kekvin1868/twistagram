package controller

import (
	"net/http"
	"twistagram/src/modules/bookmark/domain"
	"twistagram/src/modules/bookmark/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func PostBookmark(c *gin.Context) {
	var bookmark domain.Bookmark

	err := c.ShouldBindJSON(&bookmark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newBookmark *domain.Bookmark
	newBookmark, err = service.PostBookmark(&bookmark)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if newBookmark == nil {
		failed := utils.Response{
			Status:  http.StatusOK,
			Message: "bookmark exist",
			Data:    nil,
		}
		c.JSON(http.StatusOK, failed)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newBookmark,
	}

	c.JSON(http.StatusOK, res)
}

func GetBookmark(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	bookmark, err := service.GetBookmark(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *bookmark,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteBookmark(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("ID"))
	err := service.DeleteBookmark(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)

}
