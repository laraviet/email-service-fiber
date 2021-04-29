package models

type Email struct {
	Subject string  `json:"subject" validate:"required"`
	From    Account `json:"from" validate:"required"`
}
