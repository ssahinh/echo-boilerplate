package middlewares

import "github.com/labstack/echo/middleware"

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
