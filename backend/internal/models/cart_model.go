package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductName string  `json:"productName"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `gorm:"" json:"totalPrice"`
	Note        string  `gorm:"type:text" json:"note"`
	ProductID   uint    `gorm:"not null;index" json:"productId"`
	UserID      uint    `gorm:"not null;index" json:"userId"`
}

type CreateCartDTO struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	Note      string  `json:"note"`
	ProductID uint    `json:"productId"`
	UserID    uint    `json:"userId"`
}

func (c CreateCartDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Price, validation.Required, is.Float, validation.Min(0.01)),
		validation.Field(&c.Qty, validation.Required, is.Int, validation.Min(1)),
		validation.Field(&c.ProductID, validation.Required, is.Int, validation.Min(1)),
		validation.Field(&c.UserID, validation.Required, is.Int, validation.Min(1)),
	)
}
