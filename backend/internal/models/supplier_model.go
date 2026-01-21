package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Supplier struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Products  []Product `json:"products"`
}

type CreateSupplierDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

func (s CreateSupplierDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&s.LastName, validation.Required, validation.Length(1, 50)),
		validation.Field(&s.Phone, validation.Required, validation.Match(regexp.MustCompile(`^+d{1,3}d{9,10}$`))),
		validation.Field(&s.Email, validation.Required, is.Email),
		validation.Field(&s.Address, validation.Required, validation.Length(5, 100)),
	)
}
