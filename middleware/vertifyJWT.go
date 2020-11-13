//package middleware
//
//import (
//	"Usermanage/model"
//	"Usermanage/util"
//	"github.com/labstack/echo"
//)
//
//func VartifyJWT() echo.MiddlewareFunc {
//	return func(c echo.HandlerFunc) {
//		authHeader := c.Request().Header.Get("Authorization")
//		tokenString := util.GetJWTToken(authHeader)
//		rst, err := model.CheckToken(tokenString)
//		if err {
//
//		}
//	}
//}
