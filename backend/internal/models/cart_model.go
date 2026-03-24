package models

import "github.com/google/uuid"

type Cart struct {
	Base
	ProductName string    `json:"productName"`
	Price       float64   `json:"price"`
	Qty         int       `json:"qty"`
	TotalPrice  float64   `json:"totalPrice"`
	Note        string    `json:"note"`
	ProductID   uuid.UUID `json:"productId"`
	UserID      uuid.UUID `gorm:"not null;index" json:"userId"`
}
