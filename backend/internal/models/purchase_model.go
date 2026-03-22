package models

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	Base
	Code       string    `json:"code"`
	Note       string    `json:"note"`
	Total      int64     `json:"total"`
	PPN        int64     `json:"ppn"`
	GrandTotal int64     `json:"grandTotal"`
	UserID     uuid.UUID `json:"userId"`
	Date       time.Time `json:"date"`
}

type PurchaseDetail struct {
	Base
	ProductID   int     `json:"productId"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalPrice"`
	PurchaseID  int     `json:"purchaseId"`
}
