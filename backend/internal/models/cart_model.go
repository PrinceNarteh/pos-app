package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Cart struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Qty        int     `json:"qty"`
	TotalPrice float64 `json:"totalPrice"`
	Note       string  `json:"note"`
	ProductID  int     `json:"productId"`
	UserID     int     `json:"userId"`
}

type CreateCartDTO struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	Note      string  `json:"note"`
	ProductID int     `json:"productId"`
	UserID    int     `json:"userId"`
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
