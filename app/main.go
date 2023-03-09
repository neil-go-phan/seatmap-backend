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
	db := repository.GetDB()
	
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})	
	userRepo := repository.NewUserRepo(db)
	userService:= services.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	roleRepo := repository.NewRoleRepo(db)
	roleService:= services.NewRoleService(roleRepo)
	rolehandler := handler.NewRoleHandler(roleService)
	roleRoutes := routes.NewRoleRoutes(rolehandler)
	roleRoutes.Setup(r)

	r.PUT("user/update", middlewares.CheckAccessToken(), rolehandler.ValidateRole, userhandler.UpdateUser)
	return r
}



