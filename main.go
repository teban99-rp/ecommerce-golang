package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/database"
	"github.com/teban99-rp/ecommerce-golang/controllers"
	"github.com/teban99-rp/ecommerce-golang/repositories"
	"github.com/teban99-rp/ecommerce-golang/routes"
	"github.com/teban99-rp/ecommerce-golang/services"
)

func main() {
	// Initialize database
	database.Connect()

	// Initialize repositories
	repoUser := repositories.NewUserRepository()
	serviceUser := services.NewUserService(repoUser)
	controllerUser := controllers.NewUserController(serviceUser)

	repoProduct := repositories.NewProductRepository()
	serviceProduct := services.NewProductService(repoProduct)
	controllerProduct := controllers.NewProductController(serviceProduct)

	// database.SeedProducts()
	// database.SeedInventory()

	router := gin.Default()
	routes.SetupRoutes(router, controllerUser, controllerProduct)

	router.Run(":8080")
}