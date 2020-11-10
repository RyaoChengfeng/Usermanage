package controller

import "github.com/labstack/echo"

type httpError struct {
	Message string
}


func ErrorHandler(c echo.Context, code int, msg string) error {
	return c.JSON(code, httpError{msg})
}
