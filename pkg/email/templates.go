package email

type Template struct {
	Subject string `json:"subject" example:"Account Registered"`
	Body    string `json:"body" example:"<h1>Welcome to Adomate</h1>"`
	CC      string `json:"cc" example:"no-reply@adomate.com"`
}

var Templates = map[string]Template{
	// Brand-new client, new account, new company
	"registration": {
		Subject: "Welcome to Adomate! Important Next Steps",
		Body:    "",
	},
	"new-user": {
		Subject: "Welcome to Adomate! Important Next Steps",
		Body:    "",
	},
	"new-user-notification": {
		Subject: "Adomate - A new account has been added",
		Body:    "",
	},
	"invoice": {
		Subject: "Adomate - Invoice {{.ID}}",
		Body:    "",
	},
}
