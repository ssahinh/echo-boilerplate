package home

import (
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(200, "From Home, Hello World")
}
