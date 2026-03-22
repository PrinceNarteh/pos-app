package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Code               string              `json:"code"`
	BarCode            string              `json:"barCode"`
	Name               string              `json:"name"`
	Image              string              `json:"image"`
	URL                string              `json:"url"`
	Price              float64             `json:"price"`
	Qty                int                 `json:"qty"`
	CategoryID         uuid.UUID           `json:"categoryId"`
	SupplierID         uuid.UUID           `json:"supplierId"`
	Carts              []Cart              `json:"carts"`
	OrderDetails       []OrderDetail       `json:"orderDetails"`
	PurchaseDetails    []PurchaseDetail    `json:"purchaseDetails"`
	OrderReturnDetails []OrderReturnDetail `json:"orderReturnDetails"`
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	if p.Carts == nil {
		p.Carts = []Cart{}
	}
	if p.PurchaseDetails == nil {
		p.PurchaseDetails = []PurchaseDetail{}
	}
	if p.OrderDetails == nil {
		p.OrderDetails = []OrderDetail{}
	}
	if p.OrderReturnDetails == nil {
		p.OrderReturnDetails = []OrderReturnDetail{}
	}
	return
}
