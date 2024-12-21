package services

import (
	"ms-live-sells/repositories"
)

type LiveMonitorService struct {
	UserRepo     repositories.UserRepository
	ProductRepo  repositories.ProductRepository
	OrderRepo    repositories.OrderRepository
	CustomerRepo repositories.CustomerRepository
}

// NewLiveMonitorService initializes a new live monitoring service with repositories
func NewLiveMonitorService(
	userRepo repositories.UserRepository,
	productRepo repositories.ProductRepository,
	orderRepo repositories.OrderRepository,
	customerRepo repositories.CustomerRepository,
) *LiveMonitorService {
	return &LiveMonitorService{
		UserRepo:     userRepo,
		ProductRepo:  productRepo,
		OrderRepo:    orderRepo,
		CustomerRepo: customerRepo,
	}
}
