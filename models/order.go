package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Total  float64 `gorm:"not null"`
	Status string  `gorm:"not null;default:'pending'"`
	Items  []OrderItem
}
