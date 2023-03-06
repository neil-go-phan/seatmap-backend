package presenter

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	Username string `json:"username"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}