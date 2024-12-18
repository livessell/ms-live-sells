package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"ms-live-sells/repositories"
	service "ms-live-sells/services"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // allowed origins
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	// Start the repositories
	userRepo := *repositories.NewUserRepository()
	productRepo := *repositories.NewProductRepository()
	orderRepo := *repositories.NewOrderRepository()

	// Dependency injections
	liveMonitorService := service.NewLiveMonitorService(userRepo, productRepo, orderRepo)
	//liveMonitorService := service.LiveMonitorService{UserRepo: userRepo, ProductRepo: productRepo, OrderRepo: orderRepo}

	// Starts the live monitoring
	err := liveMonitorService.StartMonitoring()
	if err != nil {
		log.Fatalf("Erro ao iniciar monitoramento: %v", err)
	}

}
