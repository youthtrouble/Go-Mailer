package smtp

import (
	"net/smtp"
	"os"

	"github.com/hngi/Team-Fierce.Backend-Golang/model"
)

var (
	hostname = os.Getenv("SMTP_HOST")
	port     = os.Getenv("SMTP_PORT")
	password = os.Getenv("PASSWORD")
)

// SMTP implements the mail service interface
type SMTP struct {
	Mail model.Mail
}

// New returns a new SMTP interface
func New() *SMTP {
	return &SMTP{}
}

//GetMail returns a reference to the embedded mail struct
func (s *SMTP) GetMail() *model.Mail {
	return &s.Mail
}

// Send email with smtp
func (s *SMTP) Send() error {
	from := s.Mail.Sender.Email
	to := s.Mail.Recipient.Email
	subject := s.Mail.Subject
	body := s.Mail.Body
	msg := []byte(
		"To: " + to +
			"Subject: " + subject +
			body)
	auth := smtp.PlainAuth("", from, password, hostname)
	err := smtp.SendMail(hostname+":"+port, auth, from, []string{to}, msg)
	return err
}

func (s *SMTP) SendWithTemplate() {}
func (s *SMTP) SendMultiple()     {}
