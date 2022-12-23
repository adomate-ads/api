package models

import "time"

type BiddingStrategy struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Strategy  string    `json:"strategy" gorm:"type:varchar(128)" example:"Cost Optimize"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
