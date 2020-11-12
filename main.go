package main

import (
	"Usermanage/config"
	"Usermanage/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//注册中间件，执行前打开db，执行完后关闭db
	e.POST("/register", controller.UserRegister)
	e.POST("/login", controller.UserLogin)

	u:=e.Group("/user")
	u.Use(middleware.JWTWithConfig(config.UserJWTConfig))

	u.PUT("/:user",controller.UserUpdateUserinfo)

	//注册一个中间件，在每次使用时都查询是否为管理员，用jwt查询
	a := e.Group("/user")
	a.Use(middleware.JWTWithConfig(config.AdminJWTConfig)) //怎么直接确认admin？

	a.GET("/:user", controller.AdminFindUserinfo)
	a.DELETE("/:user", controller.AdminDeleteUser)
	a.GET("/", controller.AdminListAllUsers)
	a.PUT("/:user", controller.AdminUpdateUserinfo)

	e.Logger.Fatal(e.Start(":1323"))
}
