package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"time"
)

type JwtCustomClaims struct {
	Username   string `json:"username"`
	Permission string `json:"permission"`
	Auth  bool   `json:"auth"`
	jwt.StandardClaims
}

var (
	JWTSecret   = "secret"
	JWTDuration = time.Hour * 12
	JWTTemporary = time.Minute * 5

	AdminJWTConfig = middleware.JWTConfig{
		SigningKey: JWTSecret,
		Claims:     JwtCustomClaims{Permission: "admin", Auth: true},
	}
	UserJWTConfig = middleware.JWTConfig{
		SigningKey: JWTSecret,
		Claims:     JwtCustomClaims{Permission: "default", Auth: true},
	}
)
