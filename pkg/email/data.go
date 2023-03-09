package email

type WelcomeData struct {
	FirstName string `json:"first_name" example:"John"`
	Company   string `json:"company" example:"Adomate"`
	Domain    string `json:"domain" example:"adomate.com"`
}

type PasswordResetData struct {
	FirstName         string `json:"first_name" example:"John"`
	PasswordResetLink string `json:"password_reset_link" example:"https://adomate.com/reset-password/1234"`
}
