package middlewares

import (
	"errors"
	"net/http"
	"seatmap-backend/api/presenter"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var TOKEN_SERECT_KEY = []byte("GolenOwl2023")

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-access-token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func validateToken(tokenString string) (*presenter.JWTClaim, error) {
	claims := &presenter.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return TOKEN_SERECT_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	return claims,nil
}
