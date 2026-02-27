package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/controllers"
	"github.com/teban99-rp/ecommerce-golang/middleware"
)

func RegisterViewRoutes(
	router *gin.Engine,
	controllerUser *controllers.UserController,
	controllerProductDTO *controllers.ProductControllerDTO,
	controllerCart *controllers.CartController,
	controllerOrder *controllers.OrderController) {

	view := router.Group("/view")
	{
		view.GET("/", controllers.ShowHome)
		view.GET("/login", controllers.ShowLogin)
		view.POST("/login", controllerUser.LoginView)
		view.GET("/logout", controllerUser.LogoutView)
		view.GET("/register", controllers.ShowRegister)
		view.GET("/products", controllerProductDTO.GetProductsView)
	}

	protected := view.Group("/")
	protected.Use(middleware.JWTAuthMiddlewareView())
	{
		protected.GET("/cart/:user_id", controllerCart.GetCartView)
		protected.POST("/add_cart", controllerCart.AddToCartView)
		protected.POST("/create_order", controllerOrder.CreateOrderView)
		protected.GET("/orders/:user_id", controllerOrder.GetOrderView)
		protected.POST("/orders/payment", controllerOrder.ProcessPaymentView)
	}

	admin := view.Group("/admin")
	admin.Use(middleware.JWTAuthMiddlewareView(), middleware.AuthorizeRole("admin"))
	{
		admin.GET("/dashboard", controllers.ShowAdminDashboard)
		admin.GET("/products", controllerProductDTO.GetProductsAdminView)
		// admin.GET("/users", controllerUser.GetUsers)
		admin.POST("/products", controllerProductDTO.CreateProductView)
		admin.GET("/product/:product_id", controllerProductDTO.EditProductView)
		admin.POST("/product/:product_id", controllerProductDTO.UpdateProductView)
		admin.POST("/delete/product/:product_id", controllerProductDTO.DeleteProductView)
		admin.GET("/orders", controllerOrder.GetOrderAdminView)
		admin.POST("/orders/ship/:id", controllerOrder.ShipOrder)
		admin.POST("/orders/cancel/:id", controllerOrder.CancelOrder)
	}
}
