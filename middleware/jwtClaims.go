package middleware

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

