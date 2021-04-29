package response

import (
	"github.com/gofiber/fiber/v2"
)

func ValidationError(err error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
}
