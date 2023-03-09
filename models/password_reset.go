package models

import (
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

func (pr *PasswordReset) CreatePasswordReset() error {
	err := DB.Create(&pr).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPasswordResets() ([]PasswordReset, error) {
	var passwordResets []PasswordReset
	if err := DB.Preload("User").Find(&passwordResets).Error; err != nil {
		return nil, err
	}
	return passwordResets, nil
}

func GetPasswordReset(id uint) (*PasswordReset, error) {
	var passwordReset PasswordReset
	err := DB.Where("id = ?", id).Preload("User").First(&passwordReset).Error
	if err != nil {
		return nil, err
	}
	return &passwordReset, nil
}

func GetPasswordResetByUUID(uuid string) (*PasswordReset, error) {
	var passwordReset PasswordReset
	err := DB.Where("uuid = ?", uuid).Preload("User").First(&passwordReset).Error
	if err != nil {
		return nil, err
	}
	return &passwordReset, nil
}

func (pr *PasswordReset) DeletePasswordReset() error {
	err := DB.Delete(&pr).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *PasswordReset) Expired() bool {
	return pr.CreatedAt.Add(time.Hour * 24).Before(time.Now())
}
