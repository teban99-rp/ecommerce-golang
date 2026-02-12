package repositories

import (
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/models"
)

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}
