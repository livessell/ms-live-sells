package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ms-live-sells/config"
	"ms-live-sells/database"
	"ms-live-sells/routes"
)

func main() {
	//Starts the env config
	config.Init()

	// Start Connect
	database.ConnectDatabase()

	// Set up Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
