package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/laraviet/email-service-fiber/controllers/emailController"
)

func SetRoutes() *fiber.App {
	route := fiber.New()

	route.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("USER_BASIC_AUTH"): os.Getenv("PASSWORD_BASIC_AUTH"),
		},
	}))

	route.Post("email/send", emailController.Index)

	return route
}
