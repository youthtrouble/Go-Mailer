package controller

import (
	"github.com/hngi/Team-Fierce.Backend-Golang/smtp"

	"github.com/hngi/Team-Fierce.Backend-Golang/mailgun"
	"github.com/hngi/Team-Fierce.Backend-Golang/model"
	"github.com/hngi/Team-Fierce.Backend-Golang/sendgrid"
)

// Transport types
const (
	tpSendgrid = "sendgrid"
	tpMailgun  = "mailgun"
)

//NewMailerService returns a MailerService
//depending on what transport the user wants
func NewMailerService(transport string) model.MailerService {
	switch transport {
	case tpSendgrid:
		infoLogger.Printf("For transport %s, returning a Sendgrid mailer service", transport)
		return sendgrid.New()
	case tpMailgun:
		infoLogger.Printf("For transport %s, returning a Mailgun mailer service", transport)
		return mailgun.New()
	default:
		infoLogger.Printf("Using %s Transport\n", "SMTP")
		return smtp.New()
	}
}
