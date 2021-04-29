package services

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/laraviet/email-service-fiber/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailViaSendGrid(email models.Email) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	from := mail.NewEmail(email.From.Name, email.From.Email)
	subject := email.Subject
	to := mail.NewEmail(email.To.Name, email.To.Email)
	var plainTextContent, htmlContent string
	if email.Content.Type == "text/plain" {
		plainTextContent = email.Content.Value
	} else {
		htmlContent = email.Content.Value
	}
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
