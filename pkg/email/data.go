package email

import "time"

type WelcomeData struct {
	FirstName string `json:"user" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	Domain    string `json:"domain" example:"adomate.com"`
}

type PasswordResetData struct {
	FirstName        string `json:"first_name" example:"John"`
	PasswordResetURL string `json:"password_reset_URL" example:"https://adomate.com/reset-password/1234"`
}

type NewUser struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
}

type DeleteUser struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	Time      string `json:"time" example:"2020-01-01 00:00:00"` // Check time format
}

type NewUserNotification struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	Time      string `json:"time" example:"2020-01-01 00:00:00"`
}

type DeleteUserNotification struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	Time      string `json:"time" example:"2020-01-01 00:00:00"`
}

type DeleteCompany struct {
	// Haven't decided what is most important to send here.
	// This email will be tricky.
}

type NewInvoice struct {
	InvoiceID           uint    `json:"invoice_id" example:"1234"` // Check data type
	Company             string  `json:"company" example:"Adomate"`
	CompanyAddressLine1 string  `json:"company_address_line1" example:"Adomate"`
	CompanyAddressLine2 string  `json:"company_address_line2" example:"Adomate"`
	CompanyAddressState string  `json:"company_address_state" example:"Adomate"`
	CompanyAddressZip   int     `json:"company_address_zip" example:"Adomate"`
	PaymentMethod       string  `json:"payment_method" example:"3456"`
	PreTaxAmount        float32 `json:"pre_tax_amount" example:"1230.00"`
	TaxAmount           float32 `json:"tax_amount" example:"4.56"`
	InvoiceAmount       float64 `json:"amount" example:"1234.56"` // Check data type
	Status              string  `json:"status" example:"unpaid"`
	DueAt               string  `json:"due_at" example:"2020-01-01"` // Check time format
}

type UnpaidInvoiceReminder struct {
	InvoiceID     string  `json:"invoice_id" example:"1234"`
	Company       string  `json:"company" example:"Adomate"`
	InvoiceAmount float64 `json:"invoice_amount" example:"1234.56"`
	Domain        string  `json:"domain" example:"adomate.com"`
	DueAt         string  `json:"due_at" example:"2020-01-01"` // Check time format
}

type PaidInvoice struct {
	InvoiceID           string  `json:"invoice_id" example:"1234"`
	Company             string  `json:"company" example:"Adomate"`
	CompanyAddressLine1 string  `json:"company_address_line1" example:"Adomate"`
	CompanyAddressLine2 string  `json:"company_address_line2" example:"Adomate"`
	CompanyAddressState string  `json:"company_address_state" example:"Adomate"`
	CompanyAddressZip   int     `json:"company_address_zip" example:"Adomate"`
	Product             string  `json:"product" example:"Starter"`
	ProductPrice        float32 `json:"product_price" example:"20.00"`
	PaymentMethod       string  `json:"payment_method" example:"3456"`
	TaxAmount           float32 `json:"tax_amount" example:"1234.56"`
	InvoiceAmount       float64 `json:"invoice_amount" example:"1234.56"`
	PaidAt              string  `json:"paid_at" example:"2020-01-01"` // Check time format
}

type NewCampaign struct {
	Company   string `json:"company" example:"Adomate"`
	Campaign  string `json:"campaign" example:"Adomate Campaign"`
	StartDate string `json:"start_date" example:"2020-01-01 00:00:00"`
}

type CompleteCampaign struct {
	FirstName string    `json:"first_name" example:"John"`
	Company   string    `json:"company" example:"Adomate"`
	Campaign  string    `json:"campaign" example:"Adomate Campaign"`
	Time      time.Time `json:"time" example:"2020-01-01 00:00:00"`
	Metric1   float64   `json:"metric_1" example:"1234.56"` // Make specific
	Metric2   float64   `json:"metric_2" example:"1234.56"` // Make specific
	Metric3   float64   `json:"metric_3" example:"1234.56"` // Make specific
}

type DeleteCampaign struct {
	Company  string  `json:"company" example:"Adomate"`
	Campaign string  `json:"campaign" example:"Adomate Campaign"`
	Time     string  `json:"time" example:"2020-01-01 00:00:00"`
	Metric1  float64 `json:"metric_1" example:"1234.56"` // Make specific
	Metric2  float64 `json:"metric_2" example:"1234.56"` // Make specific
	Metric3  float64 `json:"metric_3" example:"1234.56"` // Make specific
}

type MonthlyPerformanceReport struct {
	Company  string  `json:"company" example:"Adomate"`
	Campaign string  `json:"campaign" example:"Adomate Campaign"`
	Metric1  float64 `json:"metric_1" example:"1234.56"` // Make specific
	Metric2  float64 `json:"metric_2" example:"1234.56"` // Make specific
	Metric3  float64 `json:"metric_3" example:"1234.56"` // Make specific
}

type SupportAutoResponse struct {
	SupportID   string `json:"support_id" example:"1234"`
	SupportName string `json:"support_name" example:"John"`
}

type SupportManualResponse struct {
	SupportID   string `json:"support_id" example:"1234"`
	SupportName string `json:"support_name" example:"John"`
}
