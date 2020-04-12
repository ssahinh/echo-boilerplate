package home

import (
	"ModaLast/src/helpers"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	userId := helpers.UserIDFromToken(c)
	return c.JSON(200, echo.Map{
		"userId": userId,
	})
}
