package sendgrid

import (
	"fmt"
	"log"
	"os"

	"github.com/hngi/Team-Fierce.Backend-Golang/model"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Sendgrid implements the MailService interface
type Sendgrid struct {
	Mail model.Mail
}

//New return a new Sendgrid instance
func New() *Sendgrid {
	return &Sendgrid{}
}

//GetMail returns a reference to the embedded mail struct
func (sg *Sendgrid) GetMail() *model.Mail {
	return &sg.Mail
}

//Send method from interface
func (sg *Sendgrid) Send() error {
	from := mail.NewEmail(sg.Mail.Sender.Name, sg.Mail.Sender.Email)
	subject := sg.Mail.Subject
	to := mail.NewEmail(sg.Mail.Recipient.Name, sg.Mail.Recipient.Email)
	plainTextContent := sg.Mail.Body
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	return err
}

//SendWithTemplate method
func (sg *Sendgrid) SendWithTemplate() {
	from := mail.NewEmail(sg.Mail.Sender.Name, sg.Mail.Sender.Email)
	subject := sg.Mail.Subject
	to := mail.NewEmail(sg.Mail.Recipient.Name, sg.Mail.Recipient.Email)
	plainTextContent := sg.Mail.Body
	htmlContent := sg.Mail.HTMLBody
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

//SendMultiple sends multiple emails
func (sg *Sendgrid) SendMultiple() {}
