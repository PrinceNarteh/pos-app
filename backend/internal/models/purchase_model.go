package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

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

func (p Purchase) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Code, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.Total, validation.Required, validation.Min(0)),
		validation.Field(&p.PPN, validation.Required, validation.Min(0)),
		validation.Field(&p.GrandTotal, validation.Required, validation.Min(0)),
		validation.Field(&p.UserID, validation.Required, validation.Min(1)),
	)
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

func (pd PurchaseDetail) Validate() error {
	return validation.ValidateStruct(&pd,
		validation.Field(&pd.ProductName, validation.Required, validation.Length(1, 200)),
		validation.Field(&pd.Price, validation.Required, validation.Min(0.01)),
		validation.Field(&pd.Qty, validation.Required, validation.Min(1)),
		validation.Field(&pd.TotalPrice, validation.Required, validation.Min(0.01)),
		validation.Field(&pd.ProductID, validation.Required, validation.Min(1)),
		validation.Field(&pd.PurchaseID, validation.Required, validation.Min(1)),
	)
}
