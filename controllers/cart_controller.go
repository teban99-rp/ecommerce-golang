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

// Servicios
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

// Vistas
func (c *CartController) AddToCartView(ctx *gin.Context) {

	var data dto.AddToCartDTO

	user_id := ctx.PostForm("user_id")
	product_id := ctx.PostForm("product_id")
	quantity := ctx.PostForm("quantity")

	// Convertir a los tipos correctos
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}
	productId, err := strconv.Atoi(product_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product_id inválido"})
		return
	}
	Quantity, err := strconv.Atoi(quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "quantity inválido"})
		return
	}

	data.UserID = uint(userID)
	data.ProductID = uint(productId)
	data.Quantity = Quantity

	if err := c.service.AddToCart(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creamos la cookie "flash" que dura solo 10 segundos
	ctx.SetCookie("flash_msg", "¡Producto añadido al carrito correctamente!", 10, "/", "localhost", false, true)

	// Redirigimos a la vista de productos
	ctx.Redirect(http.StatusFound, "/view/products")
}

func (c *CartController) GetCartView(ctx *gin.Context) {

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	cart, err := c.service.GetCart(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Carrito no encontrado"})
		return
	}

	total := 0.0
	for _, item := range cart {
		total += item.Total
	}

	userIDStr, err := ctx.Cookie("user_id")
	var userIDCookie uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userIDCookie = uint(id)
	}

	msg, err := ctx.Cookie("flash_msg")

	if err == nil {
		ctx.SetCookie("flash_msg", "", -1, "/", "localhost", false, true)
	}

	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Carrito",
		"view":      "cart",
		"items":     cart,
		"Total":     total,
		"user_id":   userIDCookie,
		"logged_in": userIDCookie > 0,
		"error":     msg,
	})
}
