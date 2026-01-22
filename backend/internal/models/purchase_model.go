package models

import "time"

type Purchase struct {
	ID         int       `json:"id"`
	Code       string    `json:"code"`
	Note       string    `json:"note"`
	Total      int64     `json:"total"`
	PPN        int64     `json:"ppn"`
	GrandTotal int64     `json:"grandTotal"`
	UserID     int       `json:"userId"`
	Date       time.Time `json:"date"`
}

type PurchaseDetail struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"productId"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalPrice"`
	PurchaseID  int     `json:"purchaseId"`
}
