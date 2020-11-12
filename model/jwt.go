package model

import (
	"Usermanage/config"
	"errors"
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

func ParseToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if  ok && token.Valid { // 校验token //这个校验了啥？和中间件那里的校验的区别？
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
