package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/controllers"
)

func RegisterViewRoutes(router *gin.Engine) {

	router.GET("/", controllers.ShowHome)
	router.GET("/products", controllers.ShowProducts)
	router.GET("/login", controllers.ShowLogin)
	router.GET("/register", controllers.ShowRegister)
	router.GET("/cart", controllers.ShowCart)
	router.GET("/orders", controllers.ShowOrders)
	router.GET("/admin", controllers.ShowAdminDashboard)
}
