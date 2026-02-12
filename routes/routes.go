package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/controllers"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController, productController *controllers.ProductController) {
	api := router.Group("/api")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users", userController.GetUsers)
		api.POST("/products", productController.CreateProduct)
		api.GET("/products", productController.GetProducts)
	}
}