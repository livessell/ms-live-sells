package repositories

import (
	"ms-live-sells/database"
	"ms-live-sells/models"
)

type CustomerRepository struct {
	//	db *gorm.DB
}

// NewCustomerRepository create a instance by a new CustomerRepository
//func NewCustomerRepository() *CustomerRepository {
//	return &CustomerRepository{db: database.DB}
//}

// FindByUsername find a customer by his usernmae
func (r *CustomerRepository) FindByUsername(code string) (*models.Customer, error) {
	var customer models.Customer
	err := database.DB.Where("social_network_username = ?", code).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
