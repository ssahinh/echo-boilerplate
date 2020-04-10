package main

import (
	"ModaLast/src/controllers/auth"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Database connection
	db, err := ConnectDb()
	if err != nil {
		log.Fatal("%s", err)
	}

	// Routes
	// Auth
	e.POST("/", auth.Register(db))
	e.POST("/login", auth.Login(db))

	// Start Server
	e.Logger.Fatal(e.Start(":8000"))
}
