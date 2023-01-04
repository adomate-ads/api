package models

import (
	"time"
)

type Login struct {
	UserKey   string    `json:"userkey" example:"keyName"`
	Email     string    `json:"email" example:"username@email.com"`
}

type Register struct {
	FirstName string    `json:"firstname" example:"John"`
	LastName  string    `json:"lastname" example:"Smith"`
	Email     string    `json:"email" example:"username@email.com"`
	Company   string    `json:"company" example:"CompanyName"`
}

type Logout struct {
	UserKey   string    `json:"userkey" example:"keyName"`
}
