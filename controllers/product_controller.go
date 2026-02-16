package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type ProductControllerDTO struct {
	service_dto services.ProductServiceDTO
}

func NewProductControllerDTO(service_dto services.ProductServiceDTO) *ProductControllerDTO {
	return &ProductControllerDTO{service_dto}
}

func (c *ProductControllerDTO) CreateProduct(ctx *gin.Context) {
	var input dto.ProductDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service_dto.CreateProduct(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Producto creado exitosamente"})
}

func (c *ProductControllerDTO) GetProducts(ctx *gin.Context) {
	products, err := c.service_dto.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
