package models

import "time"

type EmailTemplate struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Template  string    `json:"template" gorm:"type:varchar(1000)" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
