package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginDTO struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

func (l LoginDTO) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.UsernameOrEmail, validation.Required),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 0)),
	)
}

type RegisterUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

func (r RegisterUserDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FirstName, validation.Required, validation.Length(3, 0)),
		validation.Field(&r.LastName, validation.Required, validation.Length(3, 0)),
		validation.Field(&r.Username, validation.Required, validation.Length(3, 0)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&r.Role, validation.Required, validation.In("admin", "customer").Error("invalid role. must be either 'admin' or 'customer'")))
}
