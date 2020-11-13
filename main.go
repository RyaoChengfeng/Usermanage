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

	e.POST("/register", controller.UserRegister)
	e.POST("/login", controller.UserLogin)
	e.GET("/activate", controller.UserActivate)
	e.GET("/logout", controller.Logout)

	u := e.Group("/user")
	u.Use(middleware.JWTWithConfig(config.UserJWTConfig))

	u.PUT("/:user", controller.UserUpdateUserinfo)

	a := e.Group("/user")
	a.Use(middleware.JWTWithConfig(config.AdminJWTConfig))

	a.GET("/:user", controller.AdminFindUserinfo)
	a.DELETE("/:user", controller.AdminDeleteUser)
	a.GET("/users", controller.AdminListAllUsers)
	a.PUT("/:user", controller.AdminUpdateUserinfo)

	e.Logger.Fatal(e.Start(":1323"))
}
