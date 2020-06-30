package model

//Sender details
type sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Receipient details
type recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Mail contains mail data
type Mail struct {
	Sender    sender    `json:"sender"`
	Recipient recipient `json:"recipient"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	HTMLBody  string    `json:"htmlbody"`
}
