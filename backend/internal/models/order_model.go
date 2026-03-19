package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Base
	Code         string        `json:"code"`
	Total        int64         `json:"total"`
	PPN          int64         `json:"ppn"`
	GrandTotal   int64         `json:"grandTotal"`
	UserID       int           `json:"userId"`
	OrderDetails []OrderDetail `json:"orderDetails"`
	OrderReturns []OrderReturn `json:"orderReturns"`
}

func (o *Order) AfterFind(tx *gorm.DB) (err error) {
	if o.OrderDetails == nil {
		o.OrderDetails = []OrderDetail{}
	}
	if o.OrderReturns == nil {
		o.OrderReturns = []OrderReturn{}
	}
	return
}

type OrderDetail struct {
	Base
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalPrice"`
	Note        string  `json:"note"`
	ProductID   int     `json:"productId"`
	OrderID     int     `json:"orderId"`
}

type OrderReturn struct {
	Base
	Code               string              `json:"code"`
	Note               string              `json:"note"`
	OrderID            int                 `json:"orderId"`
	UserID             int                 `json:"userId"`
	Date               time.Time           `json:"date"`
	OrderReturnDetails []OrderReturnDetail `json:"orderReturnDetails"`
}

type OrderReturnDetail struct {
	Base
	ProductID     int     `json:"productId"`
	ProductName   string  `json:"productName"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
	TotalPrice    float64 `json:"totalPrice"`
	OrderReturnID int     `json:"orderReturnId"`
}
