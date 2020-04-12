package home

import (
	"ModaLast/src/controllers/auth"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	userId := auth.UserIDFromToken(c)
	return c.JSON(200, echo.Map{
		"userId": userId,
	})
}
