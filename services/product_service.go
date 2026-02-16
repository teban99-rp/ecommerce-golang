package services

import (
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
	"gorm.io/gorm"
)

type ProductServiceDTO interface {
	CreateProduct(product *dto.ProductDTO) error
	GetProducts() ([]dto.ProductResponseDTO, error)
}

type productServiceDTO struct{}

func NewProductServiceDTO() ProductServiceDTO {
	return &productServiceDTO{}
}

func (s *productServiceDTO) CreateProduct(product *dto.ProductDTO) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		prod := models.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}

		if err := tx.Create(&prod).Error; err != nil {
			return err
		}

		inventory := models.Inventory{
			ProductID: prod.ID,
			Stock:     product.Stock,
		}
		if err := tx.Create(&inventory).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *productServiceDTO) GetProducts() ([]dto.ProductResponseDTO, error) {
	var products []models.Product
	err := database.DB.Preload("Inventory").Find(&products).Error

	if err != nil {
		return nil, err
	}

	var response []dto.ProductResponseDTO

	for _, p := range products {
		response = append(response, dto.ProductResponseDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Inventory.Stock,
		})
	}
	return response, nil
}
