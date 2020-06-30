package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hngi/Team-Fierce.Backend-Golang/controller"
)

var port string

func init() {
	// Get serve port from environment
	// Allows for easy port change
	port = os.Getenv("EMS_PORT")
	//Default to 9000
	if port == "" {
		port = "9000"
	}
}

func main() {
	// Register routers
	router := mux.NewRouter()
	router.HandleFunc("/v1/sendmail", controller.SendMailHandler).Methods(http.MethodPost)
	router.HandleFunc("/v1/sendmailwithtemplate", controller.SendMailWithTemplateHandler).Methods(http.MethodPost)

	// Spin up basic web server with localhost:{port}
	//localhost can be omitted
	fmt.Printf("Serving on port %s...\n", port)
	http.ListenAndServe(":"+port, router)
}
