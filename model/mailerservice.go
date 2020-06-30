package model

//MailerService for all mailing services
type MailerService interface {
	// Common methods
	GetMail() *Mail // Return a reference to the embedded mail struct for updating
	Send() error
	SendWithTemplate()
	SendMultiple()
}
