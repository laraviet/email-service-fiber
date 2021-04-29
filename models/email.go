package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Email struct {
	Subject string  `json:"subject"`
	From    Account `json:"from"`
}

func (e Email) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Subject, validation.Required),
		validation.Field(&e.From),
	)
}
