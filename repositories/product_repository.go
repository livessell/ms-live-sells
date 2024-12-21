package repositories

import (
	"ms-live-sells/database"
	"ms-live-sells/models"
)

type ProductRepository struct{}

// NewProductRepository cria uma nova instância do ProductRepository
//func NewProductRepository() *ProductRepository {
//	return &ProductRepository{db: database.DB}
//}

// FindByCode busca um produto pelo código
func (r *ProductRepository) FindByCode(code string) (*models.Product, error) {
	var product models.Product
	err := database.DB.Where("code = ?", code).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
