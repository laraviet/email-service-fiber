package services

import (
	"log"
	"os"

	"github.com/laraviet/email-service-fiber/models"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func printError(err error, response *rest.Response) {
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}

func sendDynamicTemplate(email models.Email) {
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(email.From.Name, email.From.Email))
	m.SetTemplateID(email.TemplateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(email.To.Name, email.To.Email),
	}
	p.AddTos(tos...)

	for key, value := range email.DynamicTemplateData {
		p.SetDynamicTemplateData(key, value)
	}
	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	printError(err, response)
}

func sendSimpleEmail(email models.Email) {
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
	printError(err, response)
}

func SendEmailViaSendGrid(email models.Email) {
	if len(email.TemplateId) > 0 {
		sendDynamicTemplate(email)
	} else {
		sendSimpleEmail(email)
	}
}
