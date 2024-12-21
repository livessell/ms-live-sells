package repositories

import (
	"ms-live-sells/database"
	"ms-live-sells/models"
	"time"
)

type OrderRepository struct{}

// NewOrderRepository cria uma nova instância do OrderRepository
//func NewOrderRepository() *OrderRepository {
//	return &OrderRepository{db: database.DB}
//}

// Create cria uma nova ordem no banco de dados
func (r *OrderRepository) Create(user *models.User, product *models.Product, customer *models.Customer) error {
	order := models.Order{
		UserID:     user.ID,
		ProductID:  product.ID,
		CustomerID: customer.ID,
		CreatedAt:  time.Now(),
	}
	return database.DB.Create(&order).Error
}
