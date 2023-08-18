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
