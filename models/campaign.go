package models

import "time"

// TODO - Linking to AdWords, we probably need to build an google-ads package and link a UUID to here.

type Campaign struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	ResourceName string    `json:"resource_name" gorm:"type:varchar(128)" example:"Primary Monthly"`
	CompanyID    uint      `json:"company_id" gorm:"type:integer" example:"1"`
	Company      Company   `json:"company" gorm:"foreignKey:CompanyID"`
	GoogleID     uint      `json:"google_id" gorm:"type:integer" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetCampaigns() ([]Campaign, error) {
	var campaigns []Campaign
	if err := DB.Preload("Company").Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

func GetCampaign(id uint) (*Campaign, error) {
	var campaign Campaign
	if err := DB.Preload("Company").First(&campaign, id).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

func GetCampaignsByCompanyID(id uint) ([]Campaign, error) {
	var campaigns []Campaign
	if err := DB.Where("company_id = ?", id).Preload("Company").Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

func GetCampaignByGoogleID(id uint) (*Campaign, error) {
	var campaign Campaign
	if err := DB.Where("google_id = ?", id).Preload("Company").First(&campaign).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (c *Campaign) CreateCampaign() error {
	err := DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Campaign) UpdateCampaign() (*Campaign, error) {
	err := DB.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Campaign) DeleteCampaign() error {
	err := DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}
