package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/models"
)

func ShowHome(c *gin.Context) {

	var products []models.Product
	database.DB.Find(&products)

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":    "Inicio",
		"view":     "home",
		"products": products,
	})
}

func ShowProducts(c *gin.Context) {

	var products []models.Product
	database.DB.Find(&products)

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":    "Productos",
		"products": products,
		"view":     "products",
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

func ShowCart(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"title": "Carrito",
		"view":  "cart",
	})
}

func ShowOrders(c *gin.Context) {

	var orders []models.Order
	database.DB.Find(&orders)

	c.HTML(http.StatusOK, "layout", gin.H{
		"title":  "Ordenes",
		"orders": orders,
		"view":   "orders",
	})
}

func ShowAdminDashboard(c *gin.Context) {

	var users []models.User
	database.DB.Find(&users)
	c.HTML(http.StatusOK, "layout", gin.H{
		"title": "Admin Dashboard",
		"users": users,
		"view":  "admin_dashboard",
	})
}
