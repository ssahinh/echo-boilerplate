package user

import (
	"ModaLast/src/controllers/auth"
	"ModaLast/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func GetUserMe(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := auth.UserIDFromToken(c)
		var user models.User
		db.Where("id = ? ", userId).First(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"data":    user,
		})
	}
}
