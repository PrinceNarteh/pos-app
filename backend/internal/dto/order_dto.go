package dto

import (
	"time"

	"github.com/PrinceNarteh/pos/internal/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderDTO struct {
	Code         string               `json:"code"`
	Total        int64                `json:"total"`
	PPN          int64                `json:"ppn"`
	GrandTotal   int64                `json:"grandTotal"`
	UserID       int                  `json:"userId"`
	OrderDetails []models.OrderDetail `json:"orderDetails"`
	OrderReturns []models.OrderReturn `json:"orderReturns"`
}

func (o OrderDTO) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Code, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.Total, validation.Required, validation.Min(0)),
		validation.Field(&o.PPN, validation.Required, validation.Min(0)),
		validation.Field(&o.GrandTotal, validation.Required, validation.Min(0)),
		validation.Field(&o.UserID, validation.Required, validation.Min(1)),
	)
}

type OrderDetailDTO struct {
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalPrice"`
	Note        string  `json:"note"`
	ProductID   int     `json:"productId"`
	OrderID     int     `json:"orderId"`
}

func (od OrderDetailDTO) Validate() error {
	return validation.ValidateStruct(&od,
		validation.Field(&od.ProductName, validation.Required, validation.Length(1, 200)),
		validation.Field(&od.Price, validation.Required, validation.Min(0.01)),
		validation.Field(&od.Qty, validation.Required, validation.Min(1)),
		validation.Field(&od.TotalPrice, validation.Required, validation.Min(0.01)),
		validation.Field(&od.ProductID, validation.Required, validation.Min(1)),
		validation.Field(&od.OrderID, validation.Required, validation.Min(1)),
	)
}

type OrderReturnDTO struct {
	Code               string                     `json:"code"`
	Note               string                     `json:"note"`
	OrderID            int                        `json:"orderId"`
	UserID             int                        `json:"userId"`
	Date               time.Time                  `json:"date"`
	OrderReturnDetails []models.OrderReturnDetail `json:"orderReturnDetails"`
}

func (or OrderReturnDTO) Validate() error {
	return validation.ValidateStruct(&or,
		validation.Field(&or.Code, validation.Required, validation.Length(1, 100)),
		validation.Field(&or.Note, validation.Length(0, 500)),
		validation.Field(&or.OrderID, validation.Required, validation.Min(1)),
		validation.Field(&or.UserID, validation.Required, validation.Min(1)),
	)
}

type OrderReturnDetailDTO struct {
	ProductID     int     `json:"productId"`
	ProductName   string  `json:"productName"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
	TotalPrice    float64 `json:"totalPrice"`
	OrderReturnID int     `json:"orderReturnId"`
}

func (ord OrderReturnDetailDTO) Validate() error {
	return validation.ValidateStruct(&ord,
		validation.Field(&ord.ProductName, validation.Required, validation.Length(1, 200)),
		validation.Field(&ord.Price, validation.Required, validation.Min(0.01)),
		validation.Field(&ord.Qty, validation.Required, validation.Min(1)),
		validation.Field(&ord.TotalPrice, validation.Required, validation.Min(0.01)),
		validation.Field(&ord.ProductID, validation.Required, validation.Min(1)),
		validation.Field(&ord.OrderReturnID, validation.Required, validation.Min(1)),
	)
}
