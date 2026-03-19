// Package models
package models

import (
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
