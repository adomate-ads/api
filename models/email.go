package models

import "time"

type Email struct {
	ID              uint `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	CompanyID       uint `json:"company"`
	Company         Company
	EmailTemplateID uint `json:"email_template"`
	EmailTemplate   EmailTemplate
	CreatedAt       time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt       time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetEmails() ([]Email, error) {
	var emails []Email
	if err := DB.Find(&emails).Error; err != nil {
		return nil, err
	}
	return emails, nil
}

func GetEmail(id uint) (*Email, error) {
	var email Email
	if err := DB.First(&email, id).Error; err != nil {
		return nil, err
	}
	return &email, nil
}

func GetEmailsByCompanyID(id uint) ([]Email, error) {
	var emails []Email
	if err := DB.Where("company_id = ?", id).Find(&emails).Error; err != nil {
		return nil, err
	}
	return emails, nil
}

func (e *Email) CreateEmail() error {
	err := DB.Create(&e).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *Email) UpdateEmail() (*Email, error) {
	err := DB.Save(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Email) DeleteEmail() error {
	err := DB.Delete(&e).Error
	if err != nil {
		return err
	}
	return nil
}
