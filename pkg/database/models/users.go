package models

import (
	"time"
)

type Login struct {
	UserKey   string    `json:"userkey" example:"keyName"`
	Email     string    `json:"email" example:"username@email.com"`
	CreatedAt time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-09-01T00:00:00Z"`
}

type Register struct {
	FirstName string    `json:"firstname" example:"John"`
	LastName  string    `json:"lastname" example:"Smith"`
	Email     string    `json:"email" example:"username@email.com"`
	Company   string    `json:"company" example:"CompanyName"`
	CreatedAt time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-09-01T00:00:00Z"`
}

type Logout struct {
	UserKey   string    `json:"userkey" example:"keyName"`
	CreatedAt time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-09-01T00:00:00Z"`
}
