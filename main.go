package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"ms-live-sells/config"
	"ms-live-sells/provider"
)

func main() {
	//Starts the env config
	config.Init()

	// Start Connect
	//database.ConnectDatabase()

	// Set up Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up routes
	//routes.SetupRoutes(e)

	Instagram := provider.InstagramAPI{AccessToken: "IGAAR8Xny5IL5BZAFB6OF9GWnNJSkpEOGVLbmNreHdac0ZADVWcxX3J6WU5OWXQ1QTY4Tm1ndFhEb2pHZAkk4bV9BWjFlRV9wUC1LU0VxeEhtWnJTOVlFU2hoc0Y0YzF2ZAHpiVjRtb1I4SUdRSGRvZAXd4VFctTHFGT2pjeVBKQkttMAZDZD"}
	_, response, err := Instagram.CheckLive("17841401499820518")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("%v", response)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
