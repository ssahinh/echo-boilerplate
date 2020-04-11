package routes

import (
	"ModaLast/src/controllers/auth"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ApiRoutes(e *echo.Echo, db *gorm.DB) {
	group := e.Group("/api/v1")
	// Auth
	group.POST("/register", auth.Register(db))
	group.POST("/login", auth.Login(db))
}
