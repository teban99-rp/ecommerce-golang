package repositories

import "github.com/teban99-rp/ecommerce-golang/models"

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
}
