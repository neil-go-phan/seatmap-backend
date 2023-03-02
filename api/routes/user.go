package routes

import (
	"seatmap-backend/api/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler handler.UserHandler
}

func NewUserRoutes(userHandler *handler.UserHandler) *UserRoutes{
	userRoute := &UserRoutes{
		userHandler: *userHandler,
	}
	return userRoute
}

func (userRoutes *UserRoutes)Setup(r *gin.Engine) {

	authRoutes := r.Group("auth")
	{
		authRoutes.POST("sign-up", userRoutes.userHandler.SignUp)
		authRoutes.POST("sign-in", userRoutes.userHandler.SignIn)
		authRoutes.GET("users", userRoutes.userHandler.GetUsers)
	}
}