package models

import (
	"gorm.io/gorm"
)

type Supplier struct {
	Base
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Products  []Product `json:"products"`
}

func (s *Supplier) AfterFind(tx *gorm.DB) (err error) {
	if s.Products == nil {
		s.Products = []Product{}
	}
	return
}
