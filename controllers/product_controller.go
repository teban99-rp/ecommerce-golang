package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type ProductControllerDTO struct {
	service_dto services.ProductServiceDTO
}

// servicios
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

func (c *ProductControllerDTO) EditProduct(ctx *gin.Context) {

	var input dto.ProductDTO
	productID, _ := strconv.Atoi(ctx.Param("product_id"))

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service_dto.EditProduct(uint(productID)); err == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Producto no encontrado"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Producto actualizado exitosamente"})
}

func (c *ProductControllerDTO) UpdateProduct(ctx *gin.Context) {

	var input dto.ProductDTO
	productID, _ := strconv.Atoi(ctx.Param("product_id"))

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service_dto.UpdateProduct(uint(productID), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Producto actualizado exitosamente"})
}

func (c *ProductControllerDTO) DeleteProduct(ctx *gin.Context) {

	productID, _ := strconv.Atoi(ctx.Param("product_id"))

	err := c.service_dto.DeleteProduct(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Producto Eliminado"})
}

// vistas
func (c *ProductControllerDTO) GetProductsView(ctx *gin.Context) {
	products, err := c.service_dto.GetProducts()
	if err != nil {
		log.Printf("Error al obtener productos: %v", err)
		ctx.Redirect(http.StatusSeeOther, "/view/home")
		return
	}

	userIDStr, err := ctx.Cookie("user_id")
	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	role, err := ctx.Cookie("role")
	msg, err := ctx.Cookie("flash_msg")

	// 2. Si existe, la "borramos" del navegador configurando MaxAge en -1
	if err == nil {
		ctx.SetCookie("flash_msg", "", -1, "/", "localhost", false, true)
	}

	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Productos",
		"view":      "products",
		"products":  products,
		"user_id":   userID,
		"logged_in": userID > 0,
		"role":      role,
		"success":   msg,
	})
}

func (c *ProductControllerDTO) GetProductsAdminView(ctx *gin.Context) {
	products, err := c.service_dto.GetProducts()
	if err != nil {
		log.Printf("Error al obtener productos: %v", err)
		ctx.Redirect(http.StatusSeeOther, "/view/home")
		return
	}

	userIDStr, err := ctx.Cookie("user_id")
	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	role, err := ctx.Cookie("role")
	msg, err := ctx.Cookie("flash_msg")

	if err == nil {
		ctx.SetCookie("flash_msg", "", -1, "/", "localhost", false, true)
	}

	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Administrador Productos",
		"view":      "admin_products",
		"products":  products,
		"user_id":   userID,
		"logged_in": userID > 0,
		"role":      role,
		"success":   msg,
	})
}

func (c *ProductControllerDTO) CreateProductView(ctx *gin.Context) {
	var input dto.ProductDTO

	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	price := ctx.PostForm("price")
	stock := ctx.PostForm("stock")

	price_i, err := strconv.ParseFloat(price, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Price = price_i

	stock_i, err := strconv.Atoi(stock)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Stock = stock_i
	input.Name = name
	input.Description = description

	if err := c.service_dto.CreateProduct(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("flash_msg", "¡Producto creado correctamente!", 10, "/", "localhost", false, true)
	ctx.Redirect(http.StatusFound, "/view/admin/products")
}

func (c *ProductControllerDTO) EditProductView(ctx *gin.Context) {

	productID, _ := strconv.Atoi(ctx.Param("product_id"))

	product := c.service_dto.EditProduct(uint(productID))
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontro ningun producto"})
		// ctx.Redirect(http.StatusFound, "/view/admin/products")
		return
	}

	log.Printf("Producto %v", product)
	userIDStr, err := ctx.Cookie("user_id")
	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	log.Printf("Usuario %v", userID)

	// var IDProduct = product.ID
	role, err := ctx.Cookie("role")

	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Editar Producto",
		"view":      "admin_product_edit",
		"product":   product,
		"user_id":   userID,
		"logged_in": userID > 0,
		"role":      role,
	})
}

func (c *ProductControllerDTO) UpdateProductView(ctx *gin.Context) {

	var input dto.ProductDTO

	productID, _ := strconv.Atoi(ctx.Param("product_id"))
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	price := ctx.PostForm("price")
	stock := ctx.PostForm("stock")

	price_i, err := strconv.ParseFloat(price, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "precio inválido"})
		return
	}

	input.Price = price_i

	stock_i, err := strconv.Atoi(stock)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "stock inválido"})
		return
	}

	input.Stock = stock_i
	input.Name = name
	input.Description = description

	if err := c.service_dto.UpdateProduct(uint(productID), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("flash_msg", "¡Producto actualizado correctamente!", 10, "/", "localhost", false, true)

	ctx.Redirect(http.StatusFound, "/view/admin/products")
}

func (c *ProductControllerDTO) DeleteProductView(ctx *gin.Context) {

	productID, _ := strconv.Atoi(ctx.Param("product_id"))

	err := c.service_dto.DeleteProduct(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("flash_msg", "¡Producto eliminado correctamente!", 10, "/", "localhost", false, true)

	ctx.Redirect(http.StatusFound, "/view/admin/products")
}
