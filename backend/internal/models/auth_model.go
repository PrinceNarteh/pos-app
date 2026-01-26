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
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r RegisterUserDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(3, 0)),
		validation.Field(&r.Username, validation.Required, validation.Length(3, 0)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&r.Role, validation.Required, validation.In("admin", "customer")))
}
