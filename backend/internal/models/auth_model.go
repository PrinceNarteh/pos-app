package models

import validation "github.com/go-ozzo/ozzo-validation"

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
