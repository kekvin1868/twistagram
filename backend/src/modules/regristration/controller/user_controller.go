package controller

import (
	"net/http"
	"twistagram/src/modules/regristration/service"
	"twistagram/src/modules/user/domain"
	utils "twistagram/src/utils/Response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	type param struct {
		Email    string
		Password string
	}

	var parameter param
	c.BindJSON(&parameter)

	user, err := service.Login(parameter.Email, parameter.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *user,
	}

	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var newUser *domain.User
	newUser, err = service.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *newUser,
	}

	c.JSON(http.StatusOK, res)

}
