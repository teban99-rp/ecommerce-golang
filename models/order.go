package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Total  float64 `gorm:"not null"`
	Items  []OrderItem
}
