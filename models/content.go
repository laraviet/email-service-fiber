package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (c Content) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Type, validation.Required, validation.In("text/html", "text/plain").Error(`The content.type can be "text/plain" or "text/html"`)),
	)
}
