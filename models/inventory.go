package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint
	Stock     int `gorm:"not null"`
}
