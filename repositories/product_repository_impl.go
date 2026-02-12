package repositories

import (
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/models"
)

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Preload("Inventory").Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error
	return &product, err
}
