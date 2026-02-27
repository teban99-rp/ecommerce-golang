package services

import (
	"errors"

	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
	"gorm.io/gorm"
)

type ProductServiceDTO interface {
	GetProducts() ([]dto.ProductResponseDTO, error)
	CreateProduct(product *dto.ProductDTO) error
	EditProduct(productID uint) (product *dto.ProductResponseDTO)
	UpdateProduct(productID uint, data *dto.ProductDTO) error
	DeleteProduct(productID uint) error
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

func (s *productServiceDTO) EditProduct(productID uint) (data *dto.ProductResponseDTO) {

	var product models.Product

	if err := database.DB.Preload("Inventory").First(&product, productID).Error; err != nil {
		return nil
	}

	item := dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Inventory.Stock,
	}

	return &item
}

func (s *productServiceDTO) UpdateProduct(productID uint, data *dto.ProductDTO) error {

	var product models.Product

	if err := database.DB.First(&product, productID).Error; err != nil {
		return errors.New("producto no encontrado")
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.Inventory.Stock = data.Stock

	return database.DB.Save(&product).Error
}

func (s *productServiceDTO) DeleteProduct(productID uint) error {
	var product models.Product

	if err := database.DB.First(&product, productID).Error; err != nil {
		return errors.New("producto no encontrado")
	}

	return database.DB.Delete(&product).Error
}
