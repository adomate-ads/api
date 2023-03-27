package email

import "html/template"

type Template struct {
	Subject string             `json:"subject" example:"Account Registered"`
	Tmpl    *template.Template `json:"-"`
}

var Templates = map[string]Template{
	// Brand-new client, new account, new company
	"welcome": {
		Subject: "Welcome to Adomate! Important Next Steps...",
		Tmpl:    template.Must(template.ParseFiles("welcome.html")),
	},
	"forgot-password": {
		Subject: "Password Reset Request",
		Tmpl:    template.Must(template.ParseFiles("forgot-password.html")),
	},
	"new-user": {
		Subject: "Welcome, %s", // User
		Tmpl:    template.Must(template.ParseFiles("new-user.html")),
	},
	"delete-user": {
		Subject: "User Account ‘%s’ Deleted", // User
		Tmpl:    template.Must(template.ParseFiles("delete-user.html")),
	},
	"new-user-notification": {
		Subject: "User Account ‘%s’ Created", // User
		Tmpl:    template.Must(template.ParseFiles("new-user-notification.html")),
	},
	"delete-user-notification": {
		Subject: "User Account ‘%s’ Deleted", // User
		Tmpl:    template.Must(template.ParseFiles("delete-user-notification.html")),
	},
	"delete-company": {
		Subject: "", // Hardest one to figure out
		Tmpl:    template.Must(template.ParseFiles("delete-company.html")),
	},
	"new-invoice": {
		Subject: "New Invoice #%s", // InvoiceID
		Tmpl:    template.Must(template.ParseFiles("new-invoice.html")),
	},
	"unpaid-invoice-reminder": {
		Subject: "Unpaid Invoice #%s", // InvoiceID
		Tmpl:    template.Must(template.ParseFiles("unpaid-invoice-reminder.html")),
	},
	"paid-invoice": {
		Subject: "Invoice #%s Paid", // InvoiceID
		Tmpl:    template.Must(template.ParseFiles("paid-invoice.html")),
	},
	"new-campaign": {
		Subject: "‘%s’ Campaign Created", // Campaign
		Tmpl:    template.Must(template.ParseFiles("new-campaign.html")),
	},
	"campaign-completed": {
		Subject: "‘%s’ Campaign Completed", // Campaign
		Tmpl:    template.Must(template.ParseFiles("campaign-completed.html")),
	},
	"delete-campaign": {
		Subject: "‘%s’ Campaign Deleted", // Campaign
		Tmpl:    template.Must(template.ParseFiles("delete-campaign.html")),
	},
	"monthly-performance-report": {
		Subject: "‘%s’ Campaign Performance Update", // Campaign
		Tmpl:    template.Must(template.ParseFiles("monthly-performance-report.html")),
	},
	"support-auto-response": {
		Subject: "Support Request #%s", // SupportID
		Tmpl:    template.Must(template.ParseFiles("support-auto-response.html")),
	},
	"support-manual-response": {
		Subject: "Re: Support Request #%s", // SupportID
		Tmpl:    template.Must(template.ParseFiles("support-manual-response.html")),
	},
}
