package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ms-live-sells/controllers"
	"ms-live-sells/repositories"
	"ms-live-sells/services"
	"ms-live-sells/social"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // allowed origins
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	userRepo := &repositories.UserRepository{}
	socialService := &social.SocialService{}
	instagramService := &services.InstagramService{UserRepo: userRepo, SocialService: socialService}

	instagramController := &controllers.InstagramController{InstagramService: instagramService}

	e.POST("/momitoring/starts", instagramController.StartMonitoring)

}
