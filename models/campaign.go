package models

import "time"

type Campaign struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name            string    `json:"name" gorm:"type:varchar(128)" example:"Primary Monthly"`
	Company         uint      `json:"company" gorm:"type:integer" example:"1"`
	Budget          uint      `json:"budget" gorm:"type:integer" example:"1000"`
	BiddingStrategy uint      `json:"bidding_strategy" gorm:"type:integer" example:"1"`
	CreatedAt       time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt       time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
