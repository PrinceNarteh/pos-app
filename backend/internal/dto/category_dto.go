package dto

import validation "github.com/go-ozzo/ozzo-validation"

type CreateCategoryDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c CreateCategoryDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Code, validation.Required, validation.Length(1, 3)),
	)
}

type UpdateCategoryDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c UpdateCategoryDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Length(2, 50)),
		validation.Field(&c.Code, validation.Length(2, 3)),
	)
}
