package main

import (
	"ModaLast/src/database"
	"ModaLast/src/routes"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Database connection
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal("%s", err)
	}

	// Routes
	routes.ApiRoutes(e, db)

	// CronJob Test
	c := cron.New()
	c.AddFunc("* * * * *", func() { fmt.Println("Cron started") })
	c.Start()

	// Start Server
	e.Logger.Fatal(e.Start(":8000"))
}
