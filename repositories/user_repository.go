package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"ms-live-sells/database"
	"ms-live-sells/models"
)

type UserRepository struct{}

// NewUserRepository cria uma nova instância do UserRepository
//func NewUserRepository() *UserRepository {
//	return &UserRepository{db: database.DB}
//}

// GetAllUsers retorna todos os usuários
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Preload("UserSocialNetwork").
		Preload("Product").
		Preload("Order").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersWithInstagram retorna os usuários que possuem conta no Instagram.
func (r *UserRepository) GetUsersWithInstagram() ([]models.User, error) {
	var users []models.User
	err := database.DB.Preload("UserSocialNetwork").
		Preload("Product").
		Preload("Order").
		Joins("JOIN users_social_networks usn ON users.id = usn.user_id").
		Where("usn.social_network_name = ?", "instagram").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersWithInstagramByID retorna um usuário específico que possui conta no Instagram.
func (r *UserRepository) GetUsersWithInstagramByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := database.DB.Preload("UsersSocialNetwork").
		Preload("Product").
		Preload("Order").
		Joins("JOIN users_social_networks usn ON users.id = usn.user_id").
		Where("usn.social_network_name = ? AND users.id = ?", "Instagram", id).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	fmt.Println(user)
	return user, nil
}

// GetUsersBySocialNetworksNameAndUsername retorna usuários com um nome de usuário específico em redes sociais.
func (r *UserRepository) GetUsersBySocialNetworksNameAndUsername(socialNetworkName string, username string) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Product").
		Preload("Order").
		Preload("UserSocialNetwork").
		Joins("JOIN users_social_networks usn ON users.id = usn.user_id").
		Where("usn.social_network_name = ? AND usn.social_network_username = ?", socialNetworkName, username).
		Find(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
