package services

import (
	"errors"
	"math"

	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(data dto.CreateOrderDTO) error
	GetOrders(userID uint) ([]dto.OrderResponseDTO, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (s *orderService) CreateOrder(data dto.CreateOrderDTO) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {

		//Buscar carrito
		var cart models.Cart
		if err := tx.Preload("Items.Product").
			Where("user_id = ?", data.UserID).
			First(&cart).Error; err != nil {
			return errors.New("carrito no encontrado")
		}

		if len(cart.Items) == 0 {
			return errors.New("el carrito está vacío")
		}

		//Crear orden
		order := models.Order{
			UserID: data.UserID,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var total float64

		//Recorrer items del carrito
		for _, item := range cart.Items {

			var inventory models.Inventory
			if err := tx.Where("product_id = ?", item.ProductID).
				First(&inventory).Error; err != nil {
				return err
			}

			if inventory.Stock < item.Quantity {
				return errors.New("stock insuficiente para el producto")
			}

			//Crear order item
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
			}

			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}

			//Descontar stock
			inventory.Stock -= item.Quantity
			if err := tx.Save(&inventory).Error; err != nil {
				return err
			}

			total += item.Product.Price * float64(item.Quantity)
		}

		// 4️⃣ Actualizar total
		order.Total = total
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		// 5️⃣ Vaciar carrito
		if err := tx.Where("cart_id = ?", cart.ID).
			Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *orderService) GetOrders(userID uint) ([]dto.OrderResponseDTO, error) {

	var order models.Order

	err := database.DB.
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Find(&order).Error

	if err != nil {
		return nil, err
	}

	var response []dto.OrderResponseDTO

	var orderResponse dto.OrderResponseDTO
	orderResponse.ID = order.ID
	orderResponse.UserID = order.UserID
	orderResponse.TotalPrice = order.Total

	for _, item := range order.Items {
		orderItem := dto.OrderItemDTO{
			ItemOrderID: item.ID,
			ProductID:   item.ProductID,
			Name:        item.Product.Name,
			Quantity:    item.Quantity,
			Price:       item.Price,
			TotalPrice:  math.Round((item.Price * float64(item.Quantity) * 100)) / 100, // Redondear a 2 decimales
		}
		orderResponse.Items = append(orderResponse.Items, orderItem)
	}

	response = append(response, orderResponse)

	return response, nil
}
