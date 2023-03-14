package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"seatmap-backend/api/presenter"
	"seatmap-backend/services"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

var ADMIN_ROLE = "Admin"

type UserHandler struct {
	handler  services.UserService
	ESClient *elasticsearch.Client
}

func NewUserHandler(handler services.UserService, ESClient *elasticsearch.Client) *UserHandler {
	userHandler := &UserHandler{
		handler:  handler,
		ESClient: ESClient,
	}
	return userHandler
}

func (userHandler *UserHandler) CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Granted permission"})
}

func (userHandler *UserHandler) Token(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	accessToken, err := generateAccessToken(username.(string), role.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request: fail to generate access token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successful token reissue", "accessToken": accessToken})
}

func (userHandler *UserHandler) GetUsers(c *gin.Context) {
	role, _ := c.Get("role")
	if role != ADMIN_ROLE {
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
	user := newServicesUser(&inputUser)
	err := validateUsernameAndPassword(user)
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
	user := newServicesUser(&inputUser)
	err := validateSignUp(user)
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

func (userHandler *UserHandler) DeleteUser(c *gin.Context) {
	role, _ := c.Get("role")
	if role != ADMIN_ROLE {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
		return
	}
	username := c.Param("username")
	err := userHandler.handler.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Fail to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete user success"})
}

func (userHandler *UserHandler) UpdateUser(c *gin.Context) {
	roleFromRequest, _ := c.Get("role")
	if roleFromRequest != ADMIN_ROLE {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
		return
	}

	var inputUser presenter.User
	c.BindJSON(&inputUser)
	user := newServicesUser(&inputUser)
	user.Role = inputUser.Role
	err := userHandler.handler.UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Fail to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update user success"})
}

func (userHandler *UserHandler) SearchUser(c *gin.Context) {
	roleFromRequest, _ := c.Get("role")
	if roleFromRequest != ADMIN_ROLE {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
		return
	}

	var query string
	if query, _ = c.GetQuery("q"); query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no search query present"})
		return
	}
	//body := `{"query" : { "match_all" : {}" }}`
	body := fmt.Sprintf(`{"query": {"multi_match": {"query": "%s", "fields": ["full_name"]}}}`, query)

	res, err := userHandler.ESClient.Search(
		userHandler.ESClient.Search.WithContext(context.Background()),
		userHandler.ESClient.Search.WithIndex("users"),
		userHandler.ESClient.Search.WithBody(strings.NewReader(body)),
		userHandler.ESClient.Search.WithPretty(),
		)
	log.Println(res)
	log.Println(err)
	var r map[string]interface{}
	log.Println(json.NewDecoder(res.Body).Decode(&r))
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer res.Body.Close()
		if res.IsError() {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "yw"}) // ERROR_REQUEST_TO_ELASTRIC_SEARCH
			return
		}
		
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // ERROR_WHEN_PARSE_RESPONSE_BODY
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": r["hits"]})
	
}
