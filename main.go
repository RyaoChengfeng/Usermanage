package main

import (
	"Usermanage/controller"
	"Usermanage/model"
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

	//注册一个中间件，在每次使用时都查询是否为管理员，用jwt查询
	u := e.Group("/user")

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	u.Use(middleware.JWTWithConfig(config))

	u.GET("/:user", controller.AdminFindUserinfo)
	u.DELETE("/:user", controller.AdminDeleteUser)
	u.GET("/", controller.AdminListAllUsers)
	u.POST("/:user",controller.AdminUpdateUserinfo)


	e.Logger.Fatal(e.Start(":1323"))
}
