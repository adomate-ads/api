package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	FirstName string `json:"first_name" gorm:"type:varchar(128)" example:"Raaj"`
	LastName  string `json:"last_name" gorm:"type:varchar(128)" example:"Patel"`
	Email     string `json:"email" gorm:"type:varchar(128)" example:"the@raajpatel.dev"`
	Password  string `json:"password" gorm:"type:varchar(128)" example:"hashed string..."`
	RoleID    uint   `json:"role"`
	Role      Role
	CompanyID uint `json:"company"`
	Company   Company
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}
