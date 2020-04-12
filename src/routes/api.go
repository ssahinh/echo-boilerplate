package routes

import (
	"ModaLast/src/controllers/auth"
	"ModaLast/src/controllers/home"
	"ModaLast/src/controllers/post"
	"ModaLast/src/controllers/user"
	"ModaLast/src/middlewares"

	//"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ApiRoutes(e *echo.Echo, db *gorm.DB) {
	group := e.Group("/api/v1")
	// Auth
	group.POST("/auth/register", auth.Register(db))
	group.POST("/auth/login", auth.Login(db))

	// Home
	group.GET("/home", home.Hello, middlewares.IsLoggedIn)

	// User
	group.GET("/user/me", user.GetUserMe(db), middlewares.IsLoggedIn)

	// Post
	group.POST("/posts", post.CreatePost(db), middlewares.IsLoggedIn)
	group.DELETE("/posts/:id", post.DeletePost(db))
	group.GET("/posts", post.GetAllPosts(db))
	group.GET("/posts/:id", post.GetPostById(db))
}
