package models

import "github.com/google/uuid"

type Cart struct {
	Base
	ProductName string    `json:"productName"`
	Qty         int       `json:"qty"`
	Price       float64   `json:"price"`
	TotalPrice  float64   `gorm:"" json:"totalPrice"`
	Note        string    `gorm:"type:text" json:"note"`
	ProductID   uuid.UUID `gorm:"not null;index" json:"productId"`
	UserID      uuid.UUID `gorm:"not null;index" json:"userId"`
}
