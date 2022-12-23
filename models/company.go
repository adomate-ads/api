package models

import "time"

type Company struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name       string `json:"name" gorm:"type:varchar(128)" example:"Google LLC"`
	Email      string `json:"email" gorm:"type:varchar(128)" example:"the@raajpatel.dev"`
	IndustryID uint   `json:"industry"`
	Industry   Industry
	Domain     string    `json:"domain" gorm:"type:varchar(128)" example:"raajpatel.dev"`
	Budget     uint      `json:"budget" gorm:"type:integer" example:"1000"`
	CreatedAt  time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
