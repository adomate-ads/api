package email

type WelcomeData struct {
	Company      string `json:"company" example:"Adomate"`
	Domain       string `json:"domain" example:"adomate.ai"`
	CreationLink string `json:"creation_link" example:"https://app.adomate.ai/new-user/1234"`
}

type PasswordResetData struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	ResetLink string `json:"reset_link" example:"https://app.adomate.ai/reset-password/1234"`
}

type InvoiceCreatedData struct {
	Company            string `json:"company" example:"Adomate"`
	AmountDue          string `json:"amount_due" example:"18.34"`
	DueAt              string `json:"due_at" example:"2020-01-01T00:00:00Z"`
	InvoiceLink        string `json:"invoice_link" example:"https://app.adomate.ai/reset-password/1234"`
	InvoiceDescription string `json:"invoice_description" example:"August monthly service charge"`
}

type InvoicePaidData struct {
	Company            string `json:"company" example:"Adomate"`
	AmountPaid         string `json:"amount_paid" example:"18.34"`
	AmountRemaining    string `json:"amount_remaining" example:"18.34"`
	InvoiceLink        string `json:"invoice_link" example:"https://app.adomate.ai/reset-password/1234"`
	InvoiceDescription string `json:"invoice_description" example:"August monthly service charge"`
}
