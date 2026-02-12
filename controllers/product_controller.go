package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/models"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, _ := c.service.GetProducts()
	ctx.JSON(http.StatusOK, products)
}
