package models

import "time"

type Order struct {
	ID           int                        `json:"id"`
	Code         string                     `json:"code"`
	Total        int64                      `json:"total"`
	PPN          int64                      `json:"ppn"`
	GrandTotal   int64                      `json:"grandTotal"`
	UserID       int                        `json:"userId"`
	OrderDetails NullableSlice[OrderDetail] `json:"orderDetails"`
	OrderReturns NullableSlice[OrderReturn] `json:"orderReturns"`
}

type OrderDetail struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalPrice"`
	Note        string  `json:"note"`
	ProductID   int     `json:"productId"`
	OrderID     int     `json:"orderId"`
}

type OrderReturn struct {
	ID                 int                 `json:"id"`
	Code               string              `json:"code"`
	Note               string              `json:"note"`
	OrderID            int                 `json:"orderId"`
	UserID             int                 `json:"userId"`
	Date               time.Time           `json:"date"`
	OrderReturnDetails []OrderReturnDetail `json:"orderReturnDetails"`
}

type OrderReturnDetail struct {
	ID            int     `json:"id"`
	ProductID     int     `json:"productId"`
	ProductName   string  `json:"productName"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
	TotalPrice    float64 `json:"totalPrice"`
	OrderReturnID int     `json:"orderReturnId"`
}
