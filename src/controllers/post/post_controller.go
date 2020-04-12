package post

import (
	"ModaLast/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		posts := []models.Post{}
		err = db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"data":    err,
			})
		}

		// Looking like Eager loading ?!??!?
		if len(posts) > 0 {
			for i, _ := range posts {
				err := db.Debug().Model(&models.User{}).Where("id = ?",
					posts[i].UserId).Take(&posts[i].User).Error
				if err != nil {
					return c.JSON(http.StatusBadRequest, echo.Map{
						"success": false,
						"error":   err,
					})
				}
			}
		}

		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"data":    posts,
		})
	}
}

func GetPostById(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")

		var err error
		post := models.Post{}
		err = db.Debug().Model(&models.Post{}).Where("id = ?", pId).First(&post).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"data":    post,
		})
	}
}

func CreatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		user := models.User{}
	}
}
