package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Email struct {
	Subject    string  `json:"subject"`
	From       Account `json:"from"`
	To         Account `json:"to"`
	Content    Content `json:"content"`
	TemplateId string  `json:"template_id"`
}

func (e Email) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Subject, validation.Required),
		validation.Field(&e.To),
		validation.Field(&e.Content, validation.When(len(e.TemplateId) == 0)),
	)
}
