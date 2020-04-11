package routes

import (
	"ModaLast/src/controllers"
	"ModaLast/src/controllers/auth"
	"ModaLast/src/middlewares"

	//"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ApiRoutes(e *echo.Echo, db *gorm.DB) {
	group := e.Group("/api/v1")
	// Auth
	group.POST("/register", auth.Register(db))
	group.POST("/login", auth.Login(db))

	group.GET("/home", controllers.Hello, middlewares.IsLoggedIn)
}
