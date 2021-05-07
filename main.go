package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/laraviet/email-service-fiber/routes"
)

func main() {
	godotenv.Load()

	app := routes.SetRoutes()

	app.Listen(":" + os.Getenv("PORT"))
}
