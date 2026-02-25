package controllers

import (
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
