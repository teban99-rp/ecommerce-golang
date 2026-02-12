package services

import (
	"github.com/teban99-rp/ecommerce-golang/models"
	"github.com/teban99-rp/ecommerce-golang/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
	//GetProductByID(id uint) (*models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}