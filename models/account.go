package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Account struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
	)
}
