package config

import (
	"Usermanage/model"
	"github.com/labstack/echo/middleware"
	"time"
)

var (
	JWTSecret   = "secret"
	JWTDuration = time.Hour * 12
	//Md5hashSecret = []byte("hash")
	Md5hashSecret  = "hash"
	AdminJWTConfig = middleware.JWTConfig{
		SigningKey: JWTSecret,
		Claims:     model.JwtCustomClaims{Permission: "admin"},
	}
	UserJWTConfig = middleware.JWTConfig{
		SigningKey: JWTSecret,
		Claims:     model.JwtCustomClaims{Permission: "default"},
	}
)
