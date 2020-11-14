package middlewares

import (
	"Usermanage/controller"
	"Usermanage/model"
	"Usermanage/util"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func CheckUserJWTConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := util.GetJWTToken(authHeader)
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return controller.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
		if claims.Permission != "default" || claims.Auth != true {
			return controller.ErrorHandler(c, http.StatusUnauthorized, "you not have the permission!")
		}
		return next(c)
	}
}

func CheckAdminJWTConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := util.GetJWTToken(authHeader)
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return controller.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
		if claims.Permission != "admin" || claims.Auth != true {
			return controller.ErrorHandler(c, http.StatusUnauthorized, "you not have the permission!")
		}
		return next(c)
	}
}

func CheckValidJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := util.GetJWTToken(authHeader)
		rst, err := model.CheckToken(tokenString)
		if rst != 1 || err != nil {
			return controller.ErrorHandler(c, http.StatusUnauthorized, "The JWTToken is expired")
		}
		return next(c)
	}
}

//HandlerFunc和MiddlewareFunc的区别
// next（c）做了一个什么事情？
