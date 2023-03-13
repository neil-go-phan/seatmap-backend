package middlewares

import (
	"errors"
	"seatmap-backend/api/handler"

	"github.com/gin-gonic/gin"
)


func ExpiredAccessTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-refresh-token")
		if tokenString == "" {
			c.Error(handler.NewErrorReponse(errors.New("refresh token string empty"), "No refresh token"))
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			c.Error(handler.NewErrorReponse(err, "Validate tolen fail"))
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}