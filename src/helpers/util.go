package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func UserIDFromToken(c echo.Context) float64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(float64)
}
