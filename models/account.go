package models

type Account struct {
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}
