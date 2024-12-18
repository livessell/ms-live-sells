package repositories

import (
	"gorm.io/gorm"
	"ms-live-sells/database"
	"ms-live-sells/models"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository cria uma nova instância do OrderRepository
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{db: database.DB}
}

// Create cria uma nova ordem no banco de dados
func (r *OrderRepository) Create(username string, product *models.Product) error {
	order := models.Order{
		Username:    username,
		ProductID:   product.ID,
		ProductName: product.Name,
		CreatedAt:   time.Now(),
	}
	return r.db.Create(&order).Error
}
