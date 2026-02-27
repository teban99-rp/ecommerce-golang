package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {

	var input dto.CreateOrderDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateOrder(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Orden creada correctamente"})
}

func (c *OrderController) GetOrders(ctx *gin.Context) {

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	orders, err := c.service.GetOrders(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron órdenes"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) ProcessPayment(ctx *gin.Context) {

	var input dto.PaymentDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.ProcessPayment(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pago aprobado"})
}

func (c *OrderController) ShipOrder(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.ShipOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Orden enviada"})
}

func (c *OrderController) CancelOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.CancelOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Orden cancelada"})
}

func (c *OrderController) GetOrderAdmin(ctx *gin.Context) {

	orders, err := c.service.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron órdenes"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// Vistas
func (c *OrderController) CreateOrderView(ctx *gin.Context) {

	var input dto.CreateOrderDTO

	user_id := ctx.PostForm("user_id")

	// Convertir a uint
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}
	input.UserID = uint(userID)

	if err := c.service.CreateOrder(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/view/orders/"+user_id)
}

func (c *OrderController) GetOrderView(ctx *gin.Context) {

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	orders, err := c.service.GetOrders(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron órdenes"})
		return
	}

	var totalOrder = 0.0
	var orderID = 0
	for _, order := range orders {
		if order.Status == "pending" {
			orderID = int(order.ID)
			totalOrder = order.TotalPrice
		}
	}
	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":          "Mis Órdenes",
		"view":           "orders",
		"orders":         orders,
		"user_id":        userID,
		"logged_in":      userID > 0,
		"lastOrderTotal": totalOrder,
		"lastOrderID":    orderID,
	})
}

func (c *OrderController) ProcessPaymentView(ctx *gin.Context) {

	var input dto.PaymentDTO

	order_id := ctx.PostForm("order_id")
	method := ctx.PostForm("method")

	// Convertir a uint
	orderID, err := strconv.Atoi(order_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "order_id inválido"})
		return
	}
	input.OrderID = uint(orderID)
	input.Method = method

	user_id, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id no encontrado en cookies"})
		return
	}

	if err := c.service.ProcessPayment(input); err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.SetCookie("error_msg", err.Error(), 10, "/", "localhost", false, true)
		ctx.Redirect(http.StatusNotFound, "/view/orders/"+user_id)
		return
	}

	ctx.Redirect(http.StatusFound, "/view/orders/"+user_id)
}

func (c *OrderController) ShipOrderView(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.ShipOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/view/admin/orders")
}

func (c *OrderController) CancelOrderView(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.CancelOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/view/admin/orders")
}

func (c *OrderController) GetOrderAdminView(ctx *gin.Context) {

	orders, err := c.service.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron órdenes"})
		return
	}

	userIDStr, err := ctx.Cookie("user_id")

	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	role, err := ctx.Cookie("role")
	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Órdenes",
		"view":      "admin_orders",
		"orders":    orders,
		"user_id":   userID,
		"logged_in": userID > 0,
		"role":      role,
	})
}
