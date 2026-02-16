package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/controllers"
	"github.com/teban99-rp/ecommerce-golang/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
	productControllerDTO *controllers.ProductControllerDTO,
	cartController *controllers.CartController,
	orderController *controllers.OrderController,
) {
	api := router.Group("/api")
	{
		api.POST("/login", userController.Login)
		api.POST("/users", userController.CreateUser)
		api.GET("/products", productControllerDTO.GetProducts)
	}

	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/users", userController.GetUsers)
		protected.POST("/products", productControllerDTO.CreateProduct)
		protected.POST("/add_cart", cartController.AddToCart)
		protected.GET("/cart/:user_id", cartController.GetCart)
		protected.POST("/create_order", orderController.CreateOrder)
		protected.GET("/orders/:user_id", orderController.GetOrders)
	}
}
