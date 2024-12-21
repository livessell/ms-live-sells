package social

import (
	"fmt"
	"log"
	"ms-live-sells/repositories"
)

type SocialService struct {
	UserRepo     repositories.UserRepository
	ProductRepo  repositories.ProductRepository
	OrderRepo    repositories.OrderRepository
	CustomerRepo repositories.CustomerRepository
}

// processProductCode processes the product code, verifies the product in the database, and creates an order
func (h *SocialService) ProcessProductCode(productCode string, username string, commentUsername string, socialNetworkName string) error {
	// Check if the user exist in database by his social network username
	user, err := h.UserRepo.GetUsersBySocialNetworksNameAndUsername(socialNetworkName, username)
	if err != nil {
		return fmt.Errorf("error fetching user from the database: %v", err)
	}

	// Check if the product exists in the database by its code
	product, err := h.ProductRepo.FindByCode(productCode)
	if err != nil {
		return fmt.Errorf("error fetching product from the database: %v", err)
	}

	// Check if the customer exists in the database by its username
	customer, err := h.CustomerRepo.FindByUsername(commentUsername)
	if err != nil {
		return fmt.Errorf("error fetching product from the database: %v", err)
	}

	// Create an order for the product using the OrderRepository
	err = h.OrderRepo.Create(user, product, customer)
	if err != nil {
		return fmt.Errorf("error creating order: %v", err)
	}

	log.Printf("Order successfully created for product %s by user %s", product.Name, username)
	return nil
}
