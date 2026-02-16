package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/controllers"
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/routes"
	"github.com/teban99-rp/ecommerce-golang/services"
)

func main() {
	// Se inicializa la conexión a la base de datos
	database.Connect()

	// Se crean los servicios y controladores

	// Usuarios
	serviceUser := services.NewUserService()
	controllerUser := controllers.NewUserController(serviceUser)

	// Productos
	serviceProductDTO := services.NewProductServiceDTO()
	controllerProductDTO := controllers.NewProductControllerDTO(serviceProductDTO)

	// Carrito
	serviceCart := services.NewCartService()
	controllerCart := controllers.NewCartController(serviceCart)

	// Ordenes
	serviceOrder := services.NewOrderService()
	controllerOrder := controllers.NewOrderController(serviceOrder)

	// database.SeedProducts()
	// database.SeedInventory()

	router := gin.Default()
	routes.SetupRoutes(
		router,
		controllerUser,
		controllerProductDTO,
		controllerCart,
		controllerOrder,
	)
	router.Run(":8080")
}
