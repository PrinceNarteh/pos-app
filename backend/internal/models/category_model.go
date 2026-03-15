// Package models
package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

func (c *Category) AfterFind(tx *gorm.DB) (err error) {
	if c.Products == nil {
		c.Products = []Product{}
	}
	return
}

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateCategoryDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 100)),
	)
}
