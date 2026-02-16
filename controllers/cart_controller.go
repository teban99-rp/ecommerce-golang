package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type CartController struct {
	service services.CartService
}

func NewCartController(service services.CartService) *CartController {
	return &CartController{service}
}

func (c *CartController) AddToCart(ctx *gin.Context) {

	var data dto.AddToCartDTO

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.AddToCart(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Producto agregado al carrito"})
}

func (c *CartController) GetCart(ctx *gin.Context) {

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	cart, err := c.service.GetCart(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Carrito no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, cart)
}
