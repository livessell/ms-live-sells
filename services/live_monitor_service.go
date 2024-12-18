package service

import (
	"fmt"
	"log"
	"ms-live-sells/models"
	"ms-live-sells/provider"
	"ms-live-sells/repositories"
	"strings"
	"sync"
)

type LiveMonitorService struct {
	UserRepo    repositories.UserRepository
	ProductRepo repositories.ProductRepository
	OrderRepo   repositories.OrderRepository
}

// NewLiveMonitorService initializes a new live monitoring service with repositories
func NewLiveMonitorService(userRepo repositories.UserRepository, productRepo repositories.ProductRepository, orderRepo repositories.OrderRepository) *LiveMonitorService {
	return &LiveMonitorService{
		UserRepo:    userRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
	}
}

// StartMonitoring begins the monitoring process for all users with active lives
func (s *LiveMonitorService) StartMonitoring() error {
	// Retrieve all users from the repository
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error retrieving users: %v", err)
	}

	// Use goroutines to monitor each live stream concurrently
	var wg sync.WaitGroup

	// Iterate over all users and check if they have an active live stream
	for _, user := range users {

		//media, hasLive, err := utils.InstagramAPI{AccessToken}.CheckLive(user.ID)
		instagramService := repositories.InstagramService{Provider: provider.InstagramProvier{}}
		instagramResp, err := instagramService.GetUserMedias(user.InstagramUsername)
		if err != nil {
			return err
		}

		// Se o usuário não tiver uma live ativa, não faz nada
		//if !hasMedia {
		//	log.Println("User does not have an active live.")
		//	return nil
		//}

		//log.Println("User has an active live. Starting live monitoring...")
		wg.Add(1)
		for _, media := range instagramResp {
			// Goroutine to monitor the live stream of the user
			go func(md models.InstagramMedia) {
				defer wg.Done()
				// Monitor the live stream and handle errors
				err := s.monitorLive(md)
				if err != nil {
					log.Printf("Error monitoring live stream for user %s: %v", md.ID, err)
				}
			}(media)
		}

	}

	// Wait for all monitoring goroutines to finish
	wg.Wait()
	return nil
}

// monitorLive monitors the live stream of a specific user and processes comments
func (s *LiveMonitorService) monitorLive(media models.InstagramMedia) error {

	instagramService := repositories.InstagramService{Provider: provider.InstagramProvier{}}
	instagramResp, err := instagramService.GetMediaComments(media.ID)
	if err != nil {
		return err
	}

	// Processar os comentários
	var wg sync.WaitGroup
	for _, comment := range instagramResp.Data {
		wg.Add(1)
		go func(comment models.InstagramComment) {
			defer wg.Done()

			// Verificar produtos mencionados nos comentários
			productCode := extractProductCode(comment.Text)
			if productCode != "" {
				err := s.processProductCode(productCode, comment.Username)
				if err != nil {
					log.Printf("Error processing product code %s: %v", productCode, err)
				}
			}

			// Verificar menções de produtos específicos, como "XYZ"
			if strings.Contains(strings.ToLower(comment.Text), "xyz") {
				log.Printf("Product XYZ mentioned in comment by %s: %s", comment.Username, comment.Text)
			}
		}(comment)
	}

	// Esperar o processamento de todos os comentários
	wg.Wait()

	return nil
}

// extractProductCode extracts the product code from a comment
// If a product code is found (indicated by a hashtag), it returns it
func extractProductCode(comment string) string {
	// Simple logic to check for a product code in the comment (e.g., starting with "#")
	if len(comment) > 0 && comment[0] == '#' {
		return comment[1:] // Return the product code after the "#" symbol
	}
	return ""
}

// processProductCode processes the product code, verifies the product in the database, and creates an order
func (s *LiveMonitorService) processProductCode(productCode string, username string) error {
	// Check if the product exists in the database by its code
	product, err := s.ProductRepo.FindByCode(productCode)
	if err != nil {
		return fmt.Errorf("error fetching product from the database: %v", err)
	}

	// Create an order for the product using the OrderRepository
	err = s.OrderRepo.Create(username, product)
	if err != nil {
		return fmt.Errorf("error creating order: %v", err)
	}

	log.Printf("Order successfully created for product %s by user %s", product.Name, username)
	return nil
}
