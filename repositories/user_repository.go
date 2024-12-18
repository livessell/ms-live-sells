package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ms-live-sells/database"
	"ms-live-sells/models"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância do UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.DB}
}

// GetAllUsers retorna todos os usuários
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersWithInstagram retorna usuários com username do Instagram
func (r *UserRepository) GetUsersWithInstagram() ([]models.User, error) {
	var users []models.User
	err := r.db.Where("instagram_username IS NOT NULL").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// HasActiveLive verifica se o usuário tem uma live ativa
func (r *UserRepository) HasActiveLive(userID uuid.UUID) (bool, error) {
	var live models.Live

	// Busca a live ativa do usuário
	err := r.db.Where("user_id = ? AND status = ? AND start_time <= ? AND (end_time >= ? OR end_time IS NULL)",
		userID, "active", time.Now(), time.Now()).First(&live).Error

	// Se não encontrar, retornamos false
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("error checking if user has active live: %v", err)
	}

	return true, nil
}
