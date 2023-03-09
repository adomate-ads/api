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
		Body:    "kehvfeuyasikhfvbeakub",
	},
	"new-user": {
		Subject: "Welcome to Adomate! Important Next Steps",
		Body:    "",
	},
	"delete-user": {
		Subject: "Adomate - Account Deleted",
		Body:    "If you believe this was an error contact your system admin... something like that",
	},
	"new-user-notification": {
		Subject: "Adomate - A new account has been added",
		Body:    "",
	},
	"delete-user-notification": {
		Subject: "Adomate - {{.Company.Name}} User Deleted",
		Body:    "User {{.User.Name}} has been deleted.",
	},
	"delete-company": {
		Subject: "Adomate - Company Account Deleted",
		Body:    "If you believe this was an error contact Adomate Support @ ... something like that",
	},
	"new-invoice": {
		Subject: "Adomate - Invoice {{.ID}}",
		Body:    "",
	},
	"unpaid-invoice-reminder": {
		Subject: "Adomate - Invoice {{.ID}} Reminder",
		Body:    "",
	},
	"paid-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Paid!",
		Body:    "",
	},
	"delete-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Deleted",
		Body:    "",
	},
	"new-campaign": {
		Subject: "Congrats! You have successful created an Adomate Campaign",
		Body:    "",
	},
	"delete-campaign": {
		Subject: "Adomate - Campaign Deleted",
		Body:    "If you believe this was an error contact your system admin... something like that",
	},
	"forgot-password": {
		Subject: "Adomate - Password Reset",
		Body:    "",
	},
}
