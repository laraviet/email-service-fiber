package emailController

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/laraviet/email-service-fiber/models"
)

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func Index(c *fiber.Ctx) error {
	var emailObject models.Email
	c.BodyParser(&emailObject)
	validate := validator.New()

	err := validate.Struct(emailObject)
	if err != nil {
		var errors []*ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
		return c.JSON(errors)

	}

	fmt.Println(emailObject.Subject)
	fmt.Println(emailObject.From.Email)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "parse successful",
	})
}
