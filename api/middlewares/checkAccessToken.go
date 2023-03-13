package middlewares

import (
	"errors"
	"seatmap-backend/api/handler"
	"seatmap-backend/api/presenter"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var TOKEN_SERECT_KEY = []byte("GolenOwl2023")

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-access-token")
		if tokenString == "" {
			c.Error(handler.NewErrorReponse(errors.New("accest token string empty"), "Validate tolen fail"))
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

func NewErrorReponse(err error, s string) {
	panic("unimplemented")
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
