package dto

import validation "github.com/go-ozzo/ozzo-validation"

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateCategoryDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 100)),
	)
}
