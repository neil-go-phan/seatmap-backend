package handler

import (
	"errors"
	"net/http"
	"seatmap-backend/api/presenter"
	"seatmap-backend/services"

	"github.com/gin-gonic/gin"
)

var ADMIN_ROLE = "Admin"

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
		c.Error(NewErrorReponse(err, presenter.ERROR_VALIDATE_TOKEN_FAIL.ErrorName))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successful token reissue", "accessToken": accessToken})
}

func (userHandler *UserHandler) GetUsers(c *gin.Context) {
	role,_ := c.Get("role")
	if role != ADMIN_ROLE {
		c.Error(NewErrorReponse(errors.New("only admin can process"), presenter.ERROR_NO_PERMISSION.ErrorName))

		return
	}
	users, err := userHandler.handler.ListUsers()
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_BAD_REQUEST.ErrorName))
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler) SignIn(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user := newServicesUser(&inputUser)
	err := validateUsernameAndPassword(user)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_INPUT_INVALID.ErrorName))
		return
	}
	// verify user
	_, err = userHandler.handler.VerifyUser(user.Username, *user)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_SIGNIN_INCORRECT.ErrorName))
		return
	}
	// generate tokens
	fullInfoUser, _ := userHandler.handler.GetUser(user.Username)
	accessToken, err := generateAccessToken(fullInfoUser.Username, fullInfoUser.Role)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_GENERATE_TOKEN_FAIL.ErrorName))
		return
	}
	refreshToken, err := GenerateRefreshToken(fullInfoUser.Username, fullInfoUser.Role)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_GENERATE_TOKEN_FAIL.ErrorName))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign in success", "accessToken": accessToken, "refreshToken": refreshToken})
}

func (userHandler *UserHandler) SignUp(c *gin.Context) {
	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user := newServicesUser(&inputUser)
	err := validateSignUp(user)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_INPUT_INVALID.ErrorName))
		return
	}

	checkUser, err := userHandler.handler.GetUser(user.Username)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_BAD_REQUEST.ErrorName))
		return
	}
	if checkUser.Username != "" {
		c.Error(NewErrorReponse(err, presenter.ERROR_USERNAME_TAKEN.ErrorName))
		return
	}
	_, err = userHandler.handler.CreateUser(user)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_BAD_REQUEST.ErrorName))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign up success"})
}

func (userHandler *UserHandler) DeleteUser(c *gin.Context) {
	role,_ := c.Get("role")
	if role != ADMIN_ROLE {
		c.Error(NewErrorReponse(errors.New("only admin can process"), presenter.ERROR_NO_PERMISSION.ErrorName))
		return
	}
	username := c.Param("username")
	err := userHandler.handler.DeleteUser(username)
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_DELETE_FAIL.ErrorName))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete user success"})
}
// TODO: optimize validate role 
func (userHandler *UserHandler) UpdateUser(c *gin.Context) {
	roleFromRequest,_ := c.Get("role")
	if roleFromRequest != ADMIN_ROLE {
		c.Error(NewErrorReponse(errors.New("only admin can process"), presenter.ERROR_NO_PERMISSION.ErrorName))
		return
	}

	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user := newServicesUser(&inputUser)
	user.Role = inputUser.Role
	err := userHandler.handler.UpdateUser(user)
	
	if err != nil {
		c.Error(NewErrorReponse(err, presenter.ERROR_UPDATE_FAIL.ErrorName))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update user success"})
}