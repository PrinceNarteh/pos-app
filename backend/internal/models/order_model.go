package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Base
	Code         string        `json:"code"`
	Total        int64         `json:"total"`
	PPN          int64         `json:"ppn"`
	GrandTotal   int64         `json:"grandTotal"`
	UserID       uuid.UUID     `json:"userId"`
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
	ProductName string    `json:"productName"`
	Price       float64   `json:"price"`
	Qty         int       `json:"qty"`
	TotalPrice  float64   `json:"totalPrice"`
	Note        string    `json:"note"`
	ProductID   uuid.UUID `json:"productId"`
	OrderID     uuid.UUID `json:"orderId"`
}

type OrderReturn struct {
	Base
	Code               string              `json:"code"`
	Note               string              `json:"note"`
	OrderID            uuid.UUID           `json:"orderId"`
	UserID             uuid.UUID           `json:"userId"`
	Date               time.Time           `json:"date"`
	OrderReturnDetails []OrderReturnDetail `json:"orderReturnDetails"`
}

type OrderReturnDetail struct {
	Base
	ProductID     int       `json:"productId"`
	ProductName   string    `json:"productName"`
	Price         float64   `json:"price"`
	Qty           int       `json:"qty"`
	TotalPrice    float64   `json:"totalPrice"`
	OrderReturnID uuid.UUID `json:"orderReturnId"`
}
