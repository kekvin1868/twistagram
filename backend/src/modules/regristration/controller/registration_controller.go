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
		failed := utils.Response{
			Status:  http.StatusBadRequest,
			Message: "username or password inkorek",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, failed)
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
		return
	}

	if newUser == nil {
		failed := utils.Response{
			Status:  http.StatusConflict,
			Message: "Email exist",
			Data:    user,
		}
		c.JSON(http.StatusConflict, failed)
		return
	}

	res := utils.Response{
		Status:  http.StatusOK,
		Message: "user registered",
		Data:    *newUser,
	}

	c.JSON(http.StatusOK, res)

}
