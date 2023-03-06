package main

import (
	"net/http"
	"seatmap-backend/api/handler"
	middleware "seatmap-backend/api/middlewares"
	"seatmap-backend/api/routes"
	"seatmap-backend/infrastructure/repository"
	"seatmap-backend/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB()

	r := SetupRouter()
	_ = r.Run(":8080")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	db := repository.GetDB()
	
	userRepo := repository.NewUserRepo(db)
	userService:= services.NewService(userRepo)
	userhandler := handler.NewUserHandler(userService)
	
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})	
	
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	return r
}



