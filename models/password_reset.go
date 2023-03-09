package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type PasswordReset struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	UUID      string    `json:"uuid" gorm:"type:varchar(128)" example:"1234-1234-1234-1234"`
	UserID    uint      `json:"user_id" gorm:"type:integer" example:"1"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
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

func (u *User) CreateUser() error {
	err := u.HashPassword()
	if err != nil {
		return err
	}
	err = DB.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	if err := DB.Preload("Company.Industry").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersByCompanyID(id uint) ([]User, error) {
	var users []User
	err := DB.Where("company_id = ?", id).Preload("Company.Industry").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id uint) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).Preload("Company.Industry").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).Preload("Company.Industry").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser() (*User, error) {
	err := DB.Model(&u).Updates(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) DeleteUser() error {
	err := DB.Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GeneratePasswordResetToken() (string, error) {
	token, err := bcrypt.GenerateFromPassword([]byte(u.Email), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (u *User) VerifyPasswordResetToken(token string) error {
	return bcrypt.CompareHashAndPassword([]byte(token), []byte(u.Email))
}
