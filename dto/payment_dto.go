package dto

type PaymentDTO struct {
	OrderID uint   `json:"order_id" binding:"required"`
	Method  string `json:"method" binding:"required"` // card, paypal, etc
}
