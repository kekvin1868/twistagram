package controller

import (
	"net/http"
	"twistagram/src/modules/photo/domain"
	"twistagram/src/modules/photo/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func StorePhoto(c *gin.Context) {
	var photo domain.Photo

	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newPhoto *domain.Photo
	newPhoto, err = service.StorePhoto(&photo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Data:   *newPhoto,
		Status: http.StatusOK,
	}

	c.JSON(http.StatusOK, res)
}

func RetreivePhoto(c *gin.Context) {
	PostID, _ := parser.ParseID(c.Param("postid"))
	photos, err := service.RetreivePhoto(PostID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Data:   *photos,
		Status: http.StatusOK,
	}

	c.JSON(http.StatusOK, res)

}
