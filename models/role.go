package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Role      string    `json:"Role" gorm:"type:varchar(128)" example:"Administrator"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
