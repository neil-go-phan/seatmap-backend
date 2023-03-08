package main

import (
	"net/http"
	"seatmap-backend/api/handler"
	middleware "seatmap-backend/api/middlewares"
	"seatmap-backend/api/routes"
	"seatmap-backend/infrastructure/repository"
	userService "seatmap-backend/services/user"
	roleService "seatmap-backend/services/role"
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
	r.Use(middleware.Cors())
	db := repository.GetDB()
	
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})	
	userRepo := repository.NewUserRepo(db)
	userService:= userService.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	roleRepo := repository.NewRoleRepo(db)
	roleService:= roleService.NewRoleService(roleRepo)
	rolehandler := handler.NewRoleHandler(roleService)
	roleRoutes := routes.NewRoleRoutes(rolehandler)
	roleRoutes.Setup(r)

	return r
}



