package routes

import (
	"seatmap-backend/api/handler"
	"seatmap-backend/api/middlewares"

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
		authRoutes.GET("token", middlewares.ExpiredAccessTokenHandler(), userRoutes.userHandler.Token)
		authRoutes.GET("users",middlewares.CheckAccessToken(), userRoutes.userHandler.GetUsers)
	}
}