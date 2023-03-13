package main

import (
	"net/http"
	"seatmap-backend/api/handler"
	"seatmap-backend/api/middlewares"
	"seatmap-backend/api/routes"
	"seatmap-backend/infrastructure/repository"
	"seatmap-backend/services"

	// "seatmap-backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB()

	r := SetupRouter()
	_ = r.Run(":8080")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Use(middlewares.JSONAppErrorReporter())
	db := repository.GetDB()
	
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})	

	roleRepo := repository.NewRoleRepo(db)
	roleService:= services.NewRoleService(roleRepo)
	rolehandler := handler.NewRoleHandler(roleService)
	roleRoutes := routes.NewRoleRoutes(rolehandler)
	roleRoutes.Setup(r)

	userRepo := repository.NewUserRepo(db)
	userService:= services.NewUserService(userRepo, roleService)
	userhandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	return r
}