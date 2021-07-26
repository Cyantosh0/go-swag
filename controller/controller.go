package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Cyantosh0/go-swag/model"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []model.User
	err := model.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := model.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
