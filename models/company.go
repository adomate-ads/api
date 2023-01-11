package models

import (
	"time"
)

type Company struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name       string    `json:"name" gorm:"type:varchar(128)" example:"Google LLC"`
	Email      string    `json:"email" gorm:"type:varchar(128)" example:"the@raajpatel.dev"`
	IndustryID uint      `json:"industry_id" gorm:"type:integer" example:"1"`
	Industry   Industry  `json:"industry" gorm:"foreignKey:IndustryID"`
	Domain     string    `json:"domain" gorm:"type:varchar(128)" example:"raajpatel.dev"`
	Budget     uint      `json:"budget" gorm:"type:integer" example:"1000"`
	CreatedAt  time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetCompanies() ([]Company, error) {
	var companies []Company
	if err := DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func GetCompany(id uint) (*Company, error) {
	var company Company
	if err := DB.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func GetCompanyByName(name string) (*Company, error) {
	var company Company
	if err := DB.Where("name = ?", name).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func GetCompanyByEmail(email string) (*Company, error) {
	var company Company
	if err := DB.Where("email = ?", email).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *Company) CreateCompany() error {
	err := DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Company) UpdateCompany() (*Company, error) {
	err := DB.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Company) DeleteCompany() error {
	err := DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}
