package model

import (
	"Usermanage/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtCustomClaims struct {
	Username   string `json:"name"`
	Permission string `json:"admin"`
	jwt.StandardClaims
}

func CreateToken(userinfo UserInfo) (string, error) {
	claims := &JwtCustomClaims{
		Username:   userinfo.UserName,
		Permission: userinfo.Permission,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.JWTDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	return tokenString, err
}


