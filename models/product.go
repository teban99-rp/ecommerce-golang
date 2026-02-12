package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64 `gorm:"not null"`
	Inventory   Inventory  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}