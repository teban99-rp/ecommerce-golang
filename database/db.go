package database

import (
	"log"

	"github.com/teban99-rp/ecommerce-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se puede conectar a la base de datos:", err)
		return
	}

	DB = database

	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Inventory{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
		return
	}

	log.Println("Conexión a la base de datos exitosa")
}
