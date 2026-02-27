package controllers

import (
	"math"
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

// The ShowRegister function in Go renders a HTML page for user registration.
func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"title": "Registro",
		"view":  "register",
	})
}

// The `ShowAdminDashboard` function retrieves orders and cancellation data to display on the admin
// dashboard in a Go application.
func ShowAdminDashboard(c *gin.Context) {

	var orders []models.Order
	err := database.DB.Where("status NOT IN ?", []string{"pending", "cancelled"}).Find(&orders).Error

	var cancels []models.Order
	cerr := database.DB.Where("status = ?", "cancelled").Find(&cancels).Error

	CancelTotal := 0.0
	if cerr == nil {
		for _, cancel := range cancels {
			CancelTotal += cancel.Total
		}
	}

	Total := 0.0
	for _, order := range orders {
		Total += order.Total
	}

	userIDStr, err := c.Cookie("user_id")
	var userID uint
	if err == nil {
		id, _ := strconv.Atoi(userIDStr)
		userID = uint(id)
	}

	role, err := c.Cookie("role")

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":         "Admin Dashboard",
		"view":          "admin_dashboard",
		"user_id":       userID,
		"role":          role,
		"logged_in":     userID > 0,
		"count_orders":  len(orders),
		"total":         math.Round((Total * 100)) / 100,
		"count_cancels": len(cancels),
		"cancel_total":  math.Round((CancelTotal * 100)) / 100,
	})
}
