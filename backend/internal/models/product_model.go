package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code               string                           `json:"code"`
	BarCode            string                           `json:"barCode"`
	Name               string                           `json:"name"`
	Image              string                           `json:"image"`
	URL                string                           `json:"url"`
	Price              float64                          `json:"price"`
	Qty                int                              `json:"qty"`
	CategoryID         int                              `json:"categoryId"`
	SupplierID         int                              `json:"supplierId"`
	Carts              NullableSlice[Cart]              `json:"carts"`
	OrderDetails       NullableSlice[OrderDetail]       `json:"orderDetails"`
	PurchaseDetails    NullableSlice[PurchaseDetail]    `json:"purchaseDetails"`
	OrderReturnDetails NullableSlice[OrderReturnDetail] `json:"orderReturnDetails"`
}

type CreateProductDTO struct {
	Code       string  `json:"code"`
	BarCode    string  `json:"barCode"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	URL        string  `json:"url"`
	Qty        int     `json:"qty"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"categoryId"`
	SupplierID int     `json:"supplierId"`
}

func (p CreateProductDTO) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Code, validation.Required),
		validation.Field(&p.BarCode, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Image, validation.Required),
		validation.Field(&p.URL, validation.Required, is.URL),
		validation.Field(&p.Qty, validation.Required, is.Int, validation.Min(0)),
		validation.Field(&p.Price, validation.Required, is.Float, validation.Min(0.01)),
		validation.Field(&p.CategoryID, validation.Required, is.Int, validation.Min(1)),
		validation.Field(&p.SupplierID, validation.Required, is.Int, validation.Min(1)),
	)
}
