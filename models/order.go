package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"not null;default:'pending'"`
	PayMethod string  `gorm:"null"`
	Items     []OrderItem
}
