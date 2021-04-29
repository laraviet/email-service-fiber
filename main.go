package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laraviet/email-service-fiber/controllers/emailController"
)

func setRoutes() *fiber.App {
	route := fiber.New()

	route.Post("email/send", emailController.Index)

	return route
}

func main() {
	app := setRoutes()

	app.Listen(":8010")
}
