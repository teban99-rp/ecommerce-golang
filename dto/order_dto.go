package dto

type CreateOrderDTO struct {
	UserID uint `json:"user_id" binding:"required"`
}

type OrderItemDTO struct {
	ItemOrderID uint    `json:"item_order_id"`
	ProductID   uint    `json:"product_id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderResponseDTO struct {
	ID         uint           `json:"id"`
	UserID     uint           `json:"user_id"`
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"`
	PayMethod  string         `json:"pay_method"`
	Date       string         `json:"date"`
	Items      []OrderItemDTO `json:"items"`
}
