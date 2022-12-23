package models

import "time"

type Industry struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Industry  string    `json:"Industry" gorm:"type:varchar(128)" example:"Health Care"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
