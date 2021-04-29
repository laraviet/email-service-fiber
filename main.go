package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/joho/godotenv"
	"github.com/laraviet/email-service-fiber/controllers/emailController"
)

func setRoutes() *fiber.App {
	route := fiber.New()

	route.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("USER_BASIC_AUTH"): os.Getenv("PASSWORD_BASIC_AUTH"),
		},
	}))

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
