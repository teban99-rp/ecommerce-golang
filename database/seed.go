package database

import "github.com/teban99-rp/ecommerce-golang/models"

func SeedProducts() {
	products := []models.Product{
		{Name: "Laptop", Description: "A high-performance laptop", Price: 999.99},
		{Name: "Smartphone", Description: "A latest model smartphone", Price: 699.99},
		{Name: "Headphones", Description: "Noise-cancelling headphones", Price: 199.99},
	}

	DB.Create(&products)
}

func SeedInventory() {
	inventory := []models.Inventory{
		{ProductID: 1, Stock: 10},
		{ProductID: 2, Stock: 20},
		{ProductID: 3, Stock: 15},
	}
	DB.Create(&inventory)
}

func SeedUser() {
	user := models.User{
		Name:     "Usuario",
		LastName: "Administrador",
		Email:    "admin@admin.com",
		Password: "$2a$10$ik06Og9mZnbVBe/npCY8PO8/e5kldqGZUYxMfIMUONzQVrNKHTEDK",
		Role:     "admin",
	}
	DB.Create(&user)
}
