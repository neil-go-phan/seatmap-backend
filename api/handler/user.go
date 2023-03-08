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

func NewUserHandler(handler services.UserService) *UserHandler {
	userHandler := &UserHandler{
		handler: handler,
	}
	return userHandler
}

func (userHandler *UserHandler) CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Granted permission"})
}

func (userHandler *UserHandler) Token(c *gin.Context) {
	username,_ := c.Get("username")
	role,_ := c.Get("role")
	accessToken, err := generateAccessToken(username.(string), role.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request: fail to generate access token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successful token reissue", "accessToken": accessToken})
}

func (userHandler *UserHandler) GetUsers(c *gin.Context) {
	role,_ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
		return
	}
	users, err := userHandler.handler.ListUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler) SignIn(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user, err := newServicesUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	err = validateUsernameAndPassword(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	// verify user
	_, err = userHandler.handler.VerifyUser(user.Username, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Username or password is incorrect"})
		return
	}
	// generate tokens
	fullInfoUser, _ := userHandler.handler.GetUser(user.Username)
	accessToken, err := generateAccessToken(fullInfoUser.Username, fullInfoUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request: fail to generate access token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(fullInfoUser.Username, fullInfoUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request : fail to generate refresh token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign in success", "accessToken": accessToken, "refreshToken": refreshToken})
}

func (userHandler *UserHandler) SignUp(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user, err := newServicesUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong with input"})
		return
	}
	err = validateSignUp(user)
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

func (UserHandler *UserHandler) DeleteUser(c *gin.Context) {
	role,_ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
		return
	}
	username := c.Param("username")
	err := UserHandler.handler.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Fail to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete user success"})
}
