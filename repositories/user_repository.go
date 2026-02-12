package repositories

import "github.com/teban99-rp/ecommerce-golang/models"

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
}
