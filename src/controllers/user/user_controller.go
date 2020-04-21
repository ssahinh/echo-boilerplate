package user

import (
	"ModaLast/src/helpers"
	"ModaLast/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

func GetUserMe(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helpers.UserIDFromToken(c)
		var user models.User
		db.Where("id = ? ", userId).First(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"data":    user,
		})
	}
}

func UpdateUserImage(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helpers.UserIDFromToken(c)

		// Source
		// TODO Add multiple file upload support
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		defer src.Close()

		// Destination
		filePath := file.Filename
		dst, err := os.Create(filePath)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		var user models.User
		err = db.Debug().Model(&models.User{}).Where("id = ?", userId).First(&user).Error

		// Create Image Model
		image := models.Image{Url: filePath}
		err = db.Debug().Model(&models.Image{}).Create(&image).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}
		user.Image = image

		err = db.Debug().Model(&models.Image{}).Where("id = ?", image.ID).Take(&user.Image).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"data":    user,
		})
	}
}
