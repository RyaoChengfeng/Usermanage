package controller

import "github.com/labstack/echo"

type httpError struct {
	Message string `json:"message,omitempty"`
}

//错误处理到底要实现哪些东西？
func ErrorHandler(c echo.Context, code int, msg string) error {
	return c.JSON(code, httpError{msg})
}