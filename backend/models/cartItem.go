package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
