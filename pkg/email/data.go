package email

import "time"

type WelcomeData struct {
	User    string `json:"user" example:"John"`
	Company string `json:"company" example:"Adomate"`
	Domain  string `json:"domain" example:"adomate.com"`
}

type PasswordResetData struct {
	User             string `json:"user" example:"John"`
	PasswordResetURL string `json:"password_reset_URL" example:"https://adomate.com/reset-password/1234"`
}

type NewUser struct {
	User    string `json:"user" example:"John"`
	Company string `json:"company" example:"Adomate"`
}

type DeleteUser struct {
	User    string    `json:"user" example:"John"`
	Company string    `json:"company" example:"Adomate"`
	Time    time.Time `json:"time" example:"2020-01-01 00:00:00"` // Check time format
}

type NewUserNotification struct {
	User    string    `json:"user" example:"John"`
	Company string    `json:"company" example:"Adomate"`
	Time    time.Time `json:"time" example:"2020-01-01 00:00:00"`
}

type DeleteUserNotification struct {
	User    string    `json:"user" example:"John"`
	Company string    `json:"company" example:"Adomate"`
	Time    time.Time `json:"time" example:"2020-01-01 00:00:00"`
}

type DeleteCompany struct {
	// Haven't decided what is most important to send here.
	// This email will be tricky.
}

type NewInvoice struct {
	InvoiceID     string  `json:"invoice_id" example:"1234"` // Check data type
	Company       string  `json:"company" example:"Adomate"`
	InvoiceAmount float64 `json:"invoice_amount" example:"1234.56"` // Check data type
}

type UnpaidInvoiceReminder struct {
	InvoiceID     string  `json:"invoice_id" example:"1234"`
	Company       string  `json:"company" example:"Adomate"`
	InvoiceAmount float64 `json:"invoice_amount" example:"1234.56"`
}

type PaidInvoice struct {
	InvoiceID     string  `json:"invoice_id" example:"1234"`
	Company       string  `json:"company" example:"Adomate"`
	InvoiceAmount float64 `json:"invoice_amount" example:"1234.56"`
}

type NewCampaign struct {
	User     string    `json:"user" example:"John"`
	Company  string    `json:"company" example:"Adomate"`
	Campaign string    `json:"campaign" example:"Adomate Campaign"`
	Time     time.Time `json:"time" example:"2020-01-01 00:00:00"`
}

type CompleteCampaign struct {
	User     string    `json:"user" example:"John"`
	Company  string    `json:"company" example:"Adomate"`
	Campaign string    `json:"campaign" example:"Adomate Campaign"`
	Time     time.Time `json:"time" example:"2020-01-01 00:00:00"`
	Metric1  float64   `json:"metric_1" example:"1234.56"` // Make specific
	Metric2  float64   `json:"metric_2" example:"1234.56"` // Make specific
	Metric3  float64   `json:"metric_3" example:"1234.56"` // Make specific
}

type DeleteCampaign struct {
	User     string    `json:"user" example:"John"`
	Company  string    `json:"company" example:"Adomate"`
	Campaign string    `json:"campaign" example:"Adomate Campaign"`
	Time     time.Time `json:"time" example:"2020-01-01 00:00:00"`
	Metric1  float64   `json:"metric_1" example:"1234.56"` // Make specific
	Metric2  float64   `json:"metric_2" example:"1234.56"` // Make specific
	Metric3  float64   `json:"metric_3" example:"1234.56"` // Make specific
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
