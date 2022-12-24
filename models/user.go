package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

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

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// RemovePassword - Clean User data before sending it to the client
func (u *User) RemovePassword() {
	u.Password = ""
}
