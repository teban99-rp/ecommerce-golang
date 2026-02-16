package dto

type AddToCartDTO struct {
	UserID    uint `json:"user_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CartItemDTO struct {
	CartID     uint   `json:"cart_id"`
	ItemCartID uint   `json:"item_cart_id"`
	ProductID  uint   `json:"product_id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
}
