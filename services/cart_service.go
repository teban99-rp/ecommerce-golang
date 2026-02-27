package services

import (
	"errors"
	"log"

	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
	"gorm.io/gorm"
)

type CartService interface {
	AddToCart(input *dto.AddToCartDTO) error
	GetCart(userID uint) ([]dto.CartItemDTO, error)
}

type cartService struct{}

func NewCartService() CartService {
	return &cartService{}
}

func (s *cartService) AddToCart(input *dto.AddToCartDTO) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {

		var inventory models.Inventory

		if err := tx.Where("product_id = ?", input.ProductID).First(&inventory).Error; err != nil {
			log.Printf("Error al obtener inventario: %v", err)
			return err
		}

		if inventory.Stock < input.Quantity {
			return errors.New("stock insuficiente")
		}

		//En esta sección se verifica si el usuario ya tiene un carrito, si no lo tiene se crea uno nuevo
		var cart models.Cart
		err := tx.Where("user_id = ?", input.UserID).First(&cart).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				cart = models.Cart{UserID: input.UserID}
				if err := tx.Create(&cart).Error; err != nil {
					log.Printf("Error al crear carrito: %v", err)
					return err
				}
			} else {
				log.Printf("Error al obtener carrito: %v", err)
				return err
			}
		}

		// Se agrega el producto al carrito
		cartItem := models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}

		return tx.Create(&cartItem).Error
	})
}

func (s *cartService) GetCart(userID uint) ([]dto.CartItemDTO, error) {
	var cart models.Cart

	err := database.DB.
		Preload("Items.Product").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil {
		return nil, err
	}

	var response []dto.CartItemDTO

	for _, item := range cart.Items {
		response = append(response, dto.CartItemDTO{
			CartID:     cart.ID,
			ItemCartID: item.ID,
			ProductID:  item.ProductID,
			Name:       item.Product.Name,
			Quantity:   item.Quantity,
			Price:      item.Product.Price,
			Total:      float64(item.Quantity) * item.Product.Price,
		})
	}

	return response, nil
}

func (s *cartService) ClearCart(userID uint) error {
	var cart models.Cart
	err := database.DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return err
	}

	return database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
}

func (s *cartService) RemoveItem(cartItemID uint) error {
	return database.DB.Delete(&models.CartItem{}, cartItemID).Error
}
