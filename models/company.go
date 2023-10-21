package models

import (
	"time"
)

type Company struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name        string    `json:"name" gorm:"type:varchar(128)" example:"Google LLC"`
	Email       string    `json:"email" gorm:"type:varchar(128)" example:"the@raajpatel.dev"`
	Domain      string    `json:"domain" gorm:"type:varchar(128)" example:"raajpatel.dev"`
	GoogleAdsID uint      `json:"gads_id" gorm:"type:integer" example:"1"` // Google Customer ID
	StripeID    string    `json:"stripe_id" gorm:"type:varchar(128)" example:"cus_1234567890"`
	CreatedAt   time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetCompanies() ([]Company, error) {
	var companies []Company
	if err := DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	for _, company := range companies {
		company.sanitize()
	}
	return companies, nil
}

func GetCompany(id uint) (*Company, error) {
	var company Company
	if err := DB.First(&company, id).Error; err != nil {
		return nil, err
	}
	company.sanitize()
	return &company, nil
}

func GetCompanyByName(name string) (*Company, error) {
	var company Company
	if err := DB.Where("name = ?", name).First(&company).Error; err != nil {
		return nil, err
	}
	company.sanitize()
	return &company, nil
}

func GetCompanyByEmail(email string) (*Company, error) {
	var company Company
	if err := DB.Where("email = ?", email).First(&company).Error; err != nil {
		return nil, err
	}
	company.sanitize()
	return &company, nil
}

func GetCompanyByClientID(clientID int64) (*Company, error) {
	var company Company
	if err := DB.Where("gads_id = ?", clientID).First(&company).Error; err != nil {
		return nil, err
	}
	company.sanitize()
	return &company, nil
}

func GetCompanyByStripeID(stripeID string) (*Company, error) {
	var company Company
	if err := DB.Where("stripe_id = ?", stripeID).First(&company).Error; err != nil {
		return nil, err
	}
	company.sanitize()
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
	c.sanitize()
	return c, nil
}

func (c *Company) DeleteCompany() error {
	err := DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Company) sanitize() {
	c.GoogleAdsID = 0
	c.StripeID = ""
}
