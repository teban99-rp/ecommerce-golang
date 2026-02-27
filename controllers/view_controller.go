package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/dto"
	"github.com/teban99-rp/ecommerce-golang/models"
)

func ShowHome(c *gin.Context) {

	var products []models.Product
	err := database.DB.Preload("Inventory").Limit(3).Find(&products).Error

	var response []dto.ProductResponseDTO

	if err != nil {
		c.HTML(http.StatusOK, "layout", gin.H{
			"title":    "Inicio",
			"view":     "home",
			"products": response,
		})
	}

	for _, p := range products {
		response = append(response, dto.ProductResponseDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Inventory.Stock,
		})
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":    "Inicio",
		"view":     "home",
		"products": response,
	})
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"title": "Login",
		"view":  "login",
	})
}

func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"title": "Registro",
		"view":  "register",
	})
}

func ShowAdminDashboard(c *gin.Context) {

	userIDStr, err := c.Cookie("user_id")
	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	role, err := c.Cookie("role")

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":     "Admin Dashboard",
		"view":      "admin_dashboard",
		"user_id":   userID,
		"role":      role,
		"logged_in": userID > 0,
	})
}
