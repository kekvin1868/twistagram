package controller

import (
	"fmt"
	"net/http"
	"twistagram/src/modules/user/domain"
	"twistagram/src/modules/user/service"
	utils "twistagram/src/utils/Response"
	"twistagram/src/utils/parser"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	ID, _ := parser.ParseID(c.Param("id"))
	user, err := service.GetUserData(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *user,
	}

	c.JSON(http.StatusOK, res)

}
func SearchUser(c *gin.Context) {
	users, err := service.SearchUser(c.Param("key"))
	fmt.Println(users)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if users == nil {
		noData := utils.Response{
			Data:    nil,
			Message: "No user found :)",
			Status:  http.StatusNotFound,
		}
		c.JSON(http.StatusNotFound, noData)
		return
	}

	res := utils.Response{
		Status: http.StatusOK,
		Data:   *users,
	}

	c.JSON(http.StatusOK, res)
}

func EditUserData(c *gin.Context) {
	var user *domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err = service.EditUserData(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := utils.Response{
		Status: http.StatusAccepted,
		Data:   *user,
	}

	c.JSON(http.StatusAccepted, res)

}
