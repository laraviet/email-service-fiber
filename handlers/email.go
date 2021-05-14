package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laraviet/email-service-fiber/models"
	"github.com/laraviet/email-service-fiber/response"
	"github.com/laraviet/email-service-fiber/services"
)

func Index(c *fiber.Ctx) error {
	var email models.Email
	c.BodyParser(&email)

	if err := email.Validate(); err != nil {
		return response.ValidationError(err, c)
	}

	go services.SendEmailViaSendGrid(email)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "parse successful",
	})
}
