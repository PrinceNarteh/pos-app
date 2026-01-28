// Package models
package models

import validation "github.com/go-ozzo/ozzo-validation"

type Category struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Products NullableSlice[Product] `json:"products"`
}

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateCategoryDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 100)),
	)
}
