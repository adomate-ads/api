package email

type Template struct {
	Subject string `json:"subject" example:"Account Registered"`
	HTML    string `json:"html" example:"registration.html"`
}

var Templates = map[string]Template{
	// Brand-new client, new account, new company
	"registration": {
		Subject: "Welcome to Adomate! Important Next Steps",
		HTML:    "welcome.html",
	},
	"new-user": {
		Subject: "Welcome to Adomate! Important Next Steps",
		HTML:    "",
	},
	"delete-user": {
		Subject: "Adomate - Account Deleted",
		HTML:    "If you believe this was an error contact your system admin... something like that",
	},
	"new-user-notification": {
		Subject: "Adomate - A new account has been added",
		HTML:    "",
	},
	"delete-user-notification": {
		Subject: "Adomate - {{.Company.Name}} User Deleted",
		HTML:    "User {{.User.Name}} has been deleted.",
	},
	"delete-company": {
		Subject: "Adomate - Company Account Deleted",
		HTML:    "If you believe this was an error contact Adomate Support @ ... something like that",
	},
	"new-invoice": {
		Subject: "Adomate - Invoice {{.ID}}",
		HTML:    "",
	},
	"unpaid-invoice-reminder": {
		Subject: "Adomate - Invoice {{.ID}} Reminder",
		HTML:    "",
	},
	"paid-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Paid!",
		HTML:    "",
	},
	"delete-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Deleted",
		HTML:    "",
	},
	"new-campaign": {
		Subject: "Congrats! You have successful created an Adomate Campaign",
		HTML:    "",
	},
	"delete-campaign": {
		Subject: "Adomate - Campaign Deleted",
		HTML:    "If you believe this was an error contact your system admin... something like that",
	},
	"forgot-password": {
		Subject: "Adomate - Password Reset",
		HTML:    "",
	},
}
