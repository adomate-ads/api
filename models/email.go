package models

import "time"

type Email struct {
	ID          uint `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	CompanyID   uint `json:"company"`
	Company     Company
	EmailType   uint `json:"email_type"`
	EmailTypeID EmailTemplate
	CreatedAt   time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
