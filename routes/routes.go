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
		api.POST("/register", userController.CreateUser)
		api.GET("/products", productControllerDTO.GetProducts)
		// api.GET("/logout", userController.Logout)
	}

	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		//carrito
		protected.POST("/add_cart", cartController.AddToCart)
		protected.GET("/cart/:user_id", cartController.GetCart)
		//ordenes
		protected.POST("/create_order", orderController.CreateOrder)
		protected.GET("/orders/:user_id", orderController.GetOrders)
		protected.POST("/orders/payment", orderController.ProcessPayment)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.JWTAuthMiddleware(), middleware.AuthorizeRole("admin"))
	{
		//usuarios
		admin.GET("/users", userController.GetUsers)
		admin.POST("/users/change_rol/:user_id", userController.ChangeRol)
		//productos
		admin.POST("/products", productControllerDTO.CreateProduct)
		admin.GET("/product/:product_id", productControllerDTO.EditProduct)
		admin.PUT("/product/:product_id", productControllerDTO.UpdateProduct)
		admin.DELETE("/delete/product/:product_id", productControllerDTO.DeleteProduct)
		admin.PUT("/orders/ship/:id", orderController.ShipOrder)
		admin.PUT("/orders/cancel/:id", orderController.CancelOrder)

		//ordenes
		admin.GET("/orders", orderController.GetOrderAdmin)
		admin.POST("/orders/ship/:id", orderController.ShipOrder)
		admin.POST("/orders/cancel/:id", orderController.CancelOrder)
	}
}
