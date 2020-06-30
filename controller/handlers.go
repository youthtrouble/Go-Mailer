package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

// Please note that this env variable
//should be declared in this program's execution environment,
//not the environment you access the api from.
var transport = os.Getenv("MAIL_TP")

//SendMailHandler should send mail directly
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
	mService := NewMailerService(transport)
	//Get the mail struct for the service
	mail := mService.GetMail()
	//Parse request body into mail
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(mail); err != nil {
		fmt.Fprintln(w, "Invalid request body")
	}

	// Try to send mail
	if err := mService.Send(); err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, "Email sent")
}

//SendMailWithTemplateHandler should send mail with a template attached
func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
}
