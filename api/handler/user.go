package handler

import (
	"net/http"
	"seatmap-backend/api/presenter"
	"seatmap-backend/services/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	handler services.UserService
}

func NewUserHandler(handler services.UserService) *UserHandler{
	userHandler := &UserHandler{
		handler: handler,
	}
	return userHandler
}

func (userHandler *UserHandler)GetUsers(c *gin.Context) {
	users, err := userHandler.handler.ListUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler)SignIn(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user,err := NewServicesUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	err = ValidateUsernameAndPassword(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	checkUser, err := userHandler.handler.VerifyUser(user.Username, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request"})
		return
	}
	if !checkUser {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username or password is incorrect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign in success"})
}

func (userHandler *UserHandler)SignUp(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user,err := NewServicesUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	err = Validate(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}

	checkUser, err := userHandler.handler.GetUser(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request"})
		return
	}
	if checkUser.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username is already taken"})
		return
	}
	_, err = userHandler.handler.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign up success"})
}

