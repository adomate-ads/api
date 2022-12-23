package models

import "time"

type Billing struct {
	ID      uint    `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Company string  `json:"name" gorm:"type:integer" example:"1"`
	Amount  float64 `json:"amount" gorm:"type:float" example:"900.25"`
	// Available options: paid, unpaid, pending
	Status    string    `json:"status" gorm:"type:varchar(10)" example:"paid"`
	Comments  string    `json:"comments" gorm:"type:varchar(1000)" example:"Something about the invoice..."`
	DueAt     time.Time `json:"due_at" example:"2020-01-01T00:00:00Z"`
	IssuedAt  time.Time `json:"issued_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
