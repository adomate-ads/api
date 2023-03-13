package email

import "html/template"

type Template struct {
	Subject string             `json:"subject" example:"Account Registered"`
	Tmpl    *template.Template `json:"-"`
}

var Templates = map[string]Template{
	// Brand-new client, new account, new company
	"registration": {
		Subject: "Welcome to Adomate! Important Next Steps",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"new-user": {
		Subject: "Welcome to Adomate! Important Next Steps",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"delete-user": {
		Subject: "Adomate - Account Deleted",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"new-user-notification": {
		Subject: "Adomate - A new account has been added",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"delete-user-notification": {
		Subject: "Adomate - {{.Company.Name}} User Deleted",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"delete-company": {
		Subject: "Adomate - Company Account Deleted",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"new-invoice": {
		Subject: "Adomate - Invoice {{.ID}}",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"unpaid-invoice-reminder": {
		Subject: "Adomate - Invoice {{.ID}} Reminder",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"paid-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Paid!",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"delete-invoice": {
		Subject: "Adomate - Invoice {{.ID}} Deleted",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"new-campaign": {
		Subject: "Congrats! You have successful created an Adomate Campaign",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"delete-campaign": {
		Subject: "Adomate - Campaign Deleted",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
	"forgot-password": {
		Subject: "Adomate - Password Reset",
		Tmpl:    template.Must(template.ParseFiles("registration.html")),
	},
}
