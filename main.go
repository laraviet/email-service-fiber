package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/laraviet/email-service-fiber/controllers/emailController"
)

func setRoutes() *fiber.App {
	route := fiber.New()

	route.Post("email/send", emailController.Index)

	return route
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := setRoutes()

	app.Listen(":" + os.Getenv("PORT"))
}
