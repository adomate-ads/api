package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	FirstName string    `json:"first_name" gorm:"type:varchar(128)" example:"Raaj"`
	LastName  string    `json:"last_name" gorm:"type:varchar(128)" example:"Patel"`
	Email     string    `json:"email" gorm:"type:varchar(128)" example:"the@raajpatel.dev"`
	Password  string    `json:"password" gorm:"type:varchar(128)" example:"hashed string..."`
	Role      string    `json:"role" gorm:"type:varchar(128);" example:"user"`
	CompanyID uint      `json:"company_id" gorm:"type:integer" example:"1"`
	Company   Company   `json:"company" gorm:"foreignKey:CompanyID"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func VerifyPassword(password, userEmail string) (uint, error) {
	var user User
	err := DB.Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
	for _, user := range users {
		user.RemovePassword()
	}
	return users, nil
}

func GetUsersByCompanyID(id uint) ([]User, error) {
	var users []User
	err := DB.Where("company_id = ?", id).Preload("Company.Industry").Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		user.RemovePassword()
	}
	return users, nil
}

func GetUser(id uint) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).Preload("Company.Industry").First(&user).Error
	if err != nil {
		return nil, err
	}

	user.RemovePassword()
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).Preload("Company.Industry").First(&user).Error
	if err != nil {
		return nil, err
	}

	user.RemovePassword()
	return &user, nil
}

func (u *User) UpdateUser() (*User, error) {
	err := DB.Model(&u).Updates(&u).Error
	if err != nil {
		return nil, err
	}

	u.RemovePassword()
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
