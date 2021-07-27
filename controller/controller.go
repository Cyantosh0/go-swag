package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Cyantosh0/go-swag/model"
)

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} model.User
// @Failure 404 {object} object
// @Router / [get]
func GetUsers(c *gin.Context) {
	var user []model.User
	err := model.GetAllUsers(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser ... Create User
// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body model.User true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [post]
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := model.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetUserByID ... Get the user by id
// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400,404 {object} object
// @Router /{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")

	userID, err := model.StringToBinaryUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	err = model.GetUserByID(&user, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
