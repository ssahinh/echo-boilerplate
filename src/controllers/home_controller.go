package controllers

import (
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(200, "Hello World")
}
