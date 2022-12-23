package models

import "time"

type Keyword struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Keyword   string    `json:"keyword" gorm:"type:varchar(128)" example:"Dentistry"`
	CPC       float64   `json:"cpc" gorm:"type:float" example:"9.25"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
