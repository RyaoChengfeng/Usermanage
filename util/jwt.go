package util

import (
	"Usermanage/config"
	"Usermanage/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

func createToken(userinfo model.UserInfo, AuthStatus bool, duration time.Duration) (string, error) {
	claims := &config.JwtCustomClaims{
		Username:   userinfo.UserName,
		Permission: userinfo.Permission,
		Auth:       AuthStatus,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	return tokenString, err
}

func CreateAuthToken(userinfo model.UserInfo) (string, error) {
	tokenString, err := createToken(userinfo, true, config.JWTDuration)
	return tokenString, err
}

func CreateUnAuthToken(userinfo model.UserInfo) (string, error) {
	tokenString, err := createToken(userinfo, false, config.JWTTemporary)
	return tokenString, err
}

func GetJWTToken(authHeader string) string {
	parts := strings.SplitN(authHeader, " ", 2)
	tokenString := parts[1]
	return tokenString
}

func ParseToken(tokenString string) (*config.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*config.JwtCustomClaims)
	if !(ok && token.Valid) { // 校验token //这个校验了啥？和中间件那里的校验的区别？
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
